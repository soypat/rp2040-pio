// Code generated by pioasm; DO NOT EDIT.

//go:build rp2040
// +build rp2040

package main

import (
	"machine"

	pio "github.com/soypat/rp2040-pio"
)

// st7789_parallel

const st7789_parallelWrapTarget = 0
const st7789_parallelWrap = 1

var st7789_parallelProgram = pio.Program{
	Instructions: []uint16{
		//     .wrap_target
		0x6008, //  0: out    pins, 8         side 0     
		0xb042, //  1: nop                    side 1     
		//     .wrap
	},
	Origin: -1,
}

func st7789_parallelProgramDefaultConfig(offset uint8) pio.StateMachineConfig {
	cfg := pio.DefaultStateMachineConfig()
	cfg.SetWrap(offset+st7789_parallelWrapTarget, offset+st7789_parallelWrap)
	cfg.SetSideSet(1, false, false)
	return cfg;
}

// helper function to setup GPIO output and configure the SM to output on needed pins
func parallelST7789Init(sm pio.StateMachine, offset uint8, d0_pin machine.Pin, wr_pin machine.Pin) {
    d0_pin.Configure(machine.PinConfig{Mode: machine.PinPIO0})
    sm.SetConsecutivePinDirs(d0_pin, 8, true)
    cfg := st7789_parallelProgramDefaultConfig(offset)
    cfg.SetSetPins(d0_pin, 8)
    cfg.SetSidePins(wr_pin)
    cfg.SetFIFOJoin(pio.FIFO_JOIN_TX)
	cfg.SetOutShift(false, true, 8)
    maxPIOClk := uint32(32 * machine.MHz)
    sysClkHz := machine.CPUFrequency()
    clkDiv := (sysClkHz + maxPIOClk -1) / maxPIOClk
    cfg.SetClkDivIntFrac(uint16(clkDiv), 1)
	sm.Init(offset, cfg)
	sm.SetEnabled(true)
}
