//go:build tinygo
// +build tinygo

package main

import (
	"machine"

	"tinygo.org/x/drivers/ssd1289"
)

func initHardware() *ssd1289.Device {

	bus := ssd1289.NewPinBus([16]machine.Pin{
		machine.GP4,
		machine.GP5,
		machine.GP6,
		machine.GP7,
		machine.GP8,
		machine.GP9,
		machine.GP10,
		machine.GP11,
		machine.GP12,
		machine.GP13,
		machine.GP14,
		machine.GP15,
		machine.GP16,
		machine.GP17,
		machine.GP18,
		machine.GP19,
	})

	utft := ssd1289.New(machine.GP0, machine.GP1, machine.GP2, machine.GP3, bus)
	utft.Configure()
	utft.ClearDisplay()

	return &utft
}
