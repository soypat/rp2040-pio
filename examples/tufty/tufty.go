package main

import (
	"image/color"
	"machine"
	"time"

	pio "github.com/soypat/rp2040-pio"
	"tinygo.org/x/drivers"
)

const clockHz = 133000000

const (
	csPin = machine.GP4
	dcPin = machine.GP2
	wrPin = machine.GP1
	rdPin = machine.GP3
	d0Pin = machine.GP0
	blPin = machine.GP5
)

func main() {
	time.Sleep(5 * time.Second)
	println("Initializing Display")
	display := ST7789{
		cs:                csPin,
		dc:                dcPin,
		wr:                wrPin,
		rd:                rdPin,
		d0:                d0Pin,
		bl:                blPin,
		stateMachineIndex: 0,
		dmaChannel:        2,
		width:             320,
		height:            240,
		rotation:          drivers.Rotation0,
	}

	println("Initializing PIO")
	display.pio = pio.PIO0
	println("Configuring PIO")
	display.pio.Configure()
	println("Parallel Init")
	display.ParallelInit()

	// Setup DMA
	println("Setting Up DMA")
	dmaConfig := getDefaultDMAConfig(display.dmaChannel)
	setTransferDataSize(dmaConfig, DMA_SIZE_8)
	setBSwap(dmaConfig, false)
	setDREQ(dmaConfig, display.pio.Device.GetIRQ())
	dmaChannelConfigure(display.dmaChannel, dmaConfig, display.pio.Device.TXF0.Reg, 0, 0, false)

	rdPin.High()

	println("Display Common Init")
	display.CommonInit()

	println("Making Screen Blue")
	blue := color.RGBA{255, 255, 255, 255}
	display.FillRectangle(0, 0, 320, 240, blue)
}
