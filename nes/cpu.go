package nes

import "fmt"

//      - cycles
//
//

const ClockDivider byte = 12

type instruction struct {
	name           string
	size           byte
	addressingMode byte
}

// table of instructions, indexed by optcode
var instructions = [256]instruction{
	//TODO: complete with full instruction set
	0x00: {"AAA", 3, 3},
	0x01: {"ABC", 1, 1},
}

// CPU represents the 6502 CPU of the NES system
type CPU struct {
	//TODO: add all cpu internals
	regX byte
}

// Construct a new Cpu
func NewCPU() CPU {
	return CPU{}
}

// do one step
func (c *CPU) Step() {}

// run indefinitly
func (c *CPU) Run() {}

func (c *CPU) Reset()   {}
func (c *CPU) Powerup() {}

func (c *CPU) ToString() string {
	return fmt.Sprintf("CPU: %d", c.regX)
}

// implementation of instructions
func (c *CPU) Instruction_AAA() {}
