package tinyamifont

import (
	"image/color"
	"os"

	"tinygo.org/x/drivers"
)

type FontData string

type bitCursor struct {
	byteOffset int
	mask       byte
}

type Style uint8

const (
	Regular = 0
	Bold    = 1 << iota
	Italic
	Underline
	Strikethrough
)

func newBitCursor(offset int32) bitCursor {
	byteOffset := int(offset / 8)
	bitOffset := offset % 8
	mask := byte(1 << (7 - bitOffset))

	return bitCursor{
		byteOffset: byteOffset,
		mask:       mask,
	}
}

func (b *bitCursor) isBitSet(data FontData) bool {
	return data[b.byteOffset]&b.mask != 0
}

func (b *bitCursor) nextBit() {
	if b.mask == 1 {
		b.byteOffset++
		b.mask = byte(1 << 7)

	} else {
		b.mask = b.mask >> 1
	}
}

func parseUWORD(data FontData, pos int) (r uint16, next int) {
	r = uint16(data[pos+1])
	r |= uint16(data[pos]) << 8
	next = pos + 2
	return
}

func parseWORD(data FontData, pos int) (r int16, next int) {
	r = int16(data[pos+1])
	r |= int16(data[pos]) << 8
	next = pos + 2
	return
}

func parseUBYTE(data FontData, pos int) (r uint8, next int) {
	r = uint8(data[pos])
	next = pos + 1
	return
}

func parseAPTR(data FontData, pos int) (r int, next int) {
	r = int(data[pos+3])
	r |= int(data[pos+2]) << 8
	r |= int(data[pos+1]) << 16
	r |= int(data[pos]) << 24
	next = pos + 4
	return
}

type Font struct {
	Ysize         uint16
	Style         uint8
	Flags         uint8
	Xsize         uint16
	Baseline      uint16
	Boldsmear     uint16
	LoChar        uint8
	HiChar        uint8
	CharData      FontData
	Modulo        uint16
	CharLocData   FontData
	CharSpaceData FontData
	CharKernData  FontData
	Bold          bool
}

func MustLoadFont(data FontData) Font {
	font, err := LoadFont(data)
	if err != nil {
		panic(err)
	}
	return font
}

func LoadFont(data FontData) (Font, error) {
	fontHeader := data[32:] //Skip Header
	//Message struct occupies 78 bytes
	pos := 78
	tf_Ysize, pos := parseUWORD(fontHeader, pos)
	tf_Style, pos := parseUBYTE(fontHeader, pos)
	tf_Flags, pos := parseUBYTE(fontHeader, pos)
	tf_Xsize, pos := parseUWORD(fontHeader, pos)
	tf_Baseline, pos := parseUWORD(fontHeader, pos)
	tf_Boldsmear, pos := parseUWORD(fontHeader, pos)
	_, pos = parseUWORD(fontHeader, pos) //tf_Accessors
	tf_LoChar, pos := parseUBYTE(fontHeader, pos)
	tf_HighChar, pos := parseUBYTE(fontHeader, pos)
	tf_CharData, pos := parseAPTR(fontHeader, pos)
	tf_Modulo, pos := parseUWORD(fontHeader, pos)
	tf_CharLoc, pos := parseAPTR(fontHeader, pos)
	tf_CharSpace, pos := parseAPTR(fontHeader, pos)
	tf_CharKern, _ := parseAPTR(fontHeader, pos)

	/*
		println("tf_Ysize", tf_Ysize)
		println("tf_Style", tf_Style)
		println("tf_Flags", tf_Flags)
		println("tf_Xsize", tf_Xsize)
		println("tf_Baseline", tf_Baseline)
		println("tf_Boldsmear", tf_Boldsmear)
		println("tf_LoChar", tf_LoChar)
		println("tf_HighChar", tf_HighChar)
		println("tf_CharData", tf_CharData, pos)
		println("tf_Modulo", tf_Modulo)
		println("tf_CharLoc", tf_CharLoc)
		println("tf_CharSpace", tf_CharSpace)
		println("tf_CharKern", tf_CharKern)
	*/

	var charSpaceData FontData
	if tf_CharSpace != 0 {
		charSpaceData = fontHeader[tf_CharSpace:]
	}

	var charKernData FontData
	if tf_CharKern != 0 {
		charKernData = fontHeader[tf_CharKern:]
	}

	return Font{
		Ysize:         tf_Ysize,
		Style:         tf_Style,
		Flags:         tf_Flags,
		Xsize:         tf_Xsize,
		Baseline:      tf_Baseline,
		Boldsmear:     tf_Boldsmear,
		LoChar:        tf_LoChar,
		HiChar:        tf_HighChar,
		CharData:      fontHeader[tf_CharData:],
		Modulo:        tf_Modulo,
		CharLocData:   fontHeader[tf_CharLoc:],
		CharSpaceData: charSpaceData,
		CharKernData:  charKernData,
	}, nil
}

func LoadFontFile(path string) (Font, error) {
	fontFile, err := os.ReadFile(path)
	if err != nil {
		return Font{}, err
	}

	return LoadFont(FontData(fontFile))
}

func PrintString(font *Font, display drivers.Displayer, str string, x, y int16, c color.RGBA, style Style) int16 {

	underline := style&Underline != 0
	strikethrough := style&Strikethrough != 0

	len := int16(0)
	for i, chr := range str {
		len += PrintChar(font, display, uint8(chr), x+len, y, c, i != 0, style)
	}

	if underline {
		py := y + 1
		for px := int16(0); px < len; px++ {
			display.SetPixel(x+px, py, c)
		}
	}

	if strikethrough {
		py := y - (int16(font.Baseline) / 2)
		for px := int16(0); px < len; px++ {
			display.SetPixel(x+px, py, c)
		}
	}

	return len

}

func PrintChar(font *Font, display drivers.Displayer, char uint8, x, y int16, c color.RGBA, kern bool, style Style) int16 {

	if char > font.HiChar || char < font.LoChar {
		char = ' '
	}

	width, height := display.Size()

	charOffset := int(char - font.LoChar)

	kernXOffset := int16(0)
	if kern && font.CharKernData != "" {
		kernXOffset, _ = parseWORD(font.CharKernData, charOffset*2)
	}

	italicOffset := int16(0)
	if style&Italic != 0 {
		italicOffset = 1
	}

	charLocOffset := charOffset * 4
	rawBitOffset, _ := parseUWORD(font.CharLocData, charLocOffset)
	bitOffset := int32(rawBitOffset)
	bitLen, _ := parseUWORD(font.CharLocData, charLocOffset+2)
	for offsetY := uint16(0); offsetY < font.Ysize; offsetY++ {
		bitCursor := newBitCursor(bitOffset)
		if offsetY == font.Ysize/2 {
			italicOffset = 0
		}

		py := y + int16(offsetY) - int16(font.Baseline)
		if py < 0 || py >= height {
			break
		}

		for offsetX := 0; offsetX < int(bitLen); offsetX++ {
			if bitCursor.isBitSet(font.CharData) {

				px := x + int16(offsetX) + kernXOffset + italicOffset
				if px < 0 || px >= width {
					break
				}

				display.SetPixel(px, py, c)
				if style&Bold != 0 {
					if px < 0 || px >= width {
						break
					}
					px = px + int16(font.Boldsmear)
					display.SetPixel(px, py, c)
				}
			}
			bitCursor.nextBit()
		}
		bitOffset += int32(font.Modulo * 8)
	}

	if font.CharSpaceData != "" {
		glyphSize, _ := parseWORD(font.CharSpaceData, charOffset*2)
		if style&Bold != 0 {
			glyphSize += int16(font.Boldsmear)
		}
		return glyphSize + 1

	} else {
		if style&Bold != 0 {
			return int16(font.Xsize) + int16(font.Boldsmear)
		} else {
			return int16(font.Xsize)

		}
	}
}

func LineWidth(font *Font, str string, style Style) int16 {

	width := int16(0)
	for _, chr := range str {
		width += CharWidth(font, uint8(chr), style)
	}

	return width
}

func CharWidth(font *Font, char uint8, style Style) int16 {

	charOffset := int(char - font.LoChar)

	if font.CharSpaceData != "" {
		glyphSize, _ := parseWORD(font.CharSpaceData, charOffset*2)
		if style&Bold != 0 {
			glyphSize += int16(font.Boldsmear)
		}
		return glyphSize + 1

	} else {
		if style&Bold != 0 {
			return int16(font.Xsize) + int16(font.Boldsmear)
		} else {
			return int16(font.Xsize)
		}
	}
}
