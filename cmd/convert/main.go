package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const bytesPerLine = 32

var (
	fontsDir = "./fonts"
	outDir   = "."
)

func init() {
	flag.StringVar(&fontsDir, "fontsDir", fontsDir, "Directory containing amiga fonts")
	flag.StringVar(&outDir, "outDir", outDir, "directory to generate go code")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func convertFont(fileName string, outFileName string, packageName string, constName string) error {
	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	outfile, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	fmt.Fprintf(outfile, "package %s\n", packageName)
	fmt.Fprintf(outfile, "\n")
	fmt.Fprintf(outfile, "import \"github.com/spearson78/tinyamifont\"\n")
	fmt.Fprintf(outfile, "\n")
	fmt.Fprintf(outfile, "const %s =  tinyamifont.FontData(", constName)
	for i := 0; i < len(buffer); i += bytesPerLine {
		if i != 0 {
			fmt.Fprint(outfile, "+\n\t")
		}
		from, to := i, min(i+bytesPerLine, len(buffer))

		fmt.Fprintf(outfile, "%+q", buffer[from:to])
	}
	fmt.Fprintf(outfile, ")\n")

	return nil
}

func main() {
	fontDirs, err := ioutil.ReadDir(fontsDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, fontDir := range fontDirs {
		if !fontDir.IsDir() {
			continue
		}

		fontDirPath := filepath.Join(fontsDir, fontDir.Name())

		fontFiles, err := ioutil.ReadDir(fontDirPath)
		if err != nil {
			log.Fatal(err)
		}

		lowerFontName := strings.ToLower(fontDir.Name())
		outFontDir := filepath.Join(outDir, lowerFontName)

		err = os.MkdirAll(outFontDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		for _, fontFile := range fontFiles {
			if fontFile.Name() == ".fastdir" || strings.HasSuffix(fontFile.Name(), ".uaem") {
				continue
			}

			fontFilePath := filepath.Join(fontDirPath, fontFile.Name())
			outFilePath := filepath.Join(outFontDir, lowerFontName+"-"+fontFile.Name()+".go")
			constName := "Regular" + fontFile.Name() + "pt"

			//fmt.Printf("{\"%v.%v\",%v.%v},\n", lowerFontName, constName, lowerFontName, constName)

			err := convertFont(fontFilePath, outFilePath, lowerFontName, constName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
