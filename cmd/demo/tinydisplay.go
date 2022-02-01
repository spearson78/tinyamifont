//go:build !tinygo
// +build !tinygo

package main

import (
	"log"

	"github.com/sago35/tinydisplay"
)

func initHardware() *tinydisplay.Client {
	display, err := tinydisplay.NewClient("", 9812, 640, 320)
	if err != nil {
		log.Fatal(err)
	}

	return display
}
