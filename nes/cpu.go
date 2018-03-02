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
	// index: optcode
	0x00: {"AAA", 3, 3},
	0x01: {"ABC", 1, 1},
	//...
	0xFF: {"FFF", 4, 4},
}

type (
	interruptBit bool
	carryBit     bool
)

type StatusRegister byte

func (c *CPU) setCarryFlag() {
	c.statusReg = c.statusReg | (1 << 0)
}

func (c *CPU) unsetCarryFlag() {
	c.statusReg = c.statusReg | (1 << 0)
}

// CPU represents the 6502 CPU of the NES system
type CPU struct {
	//TODO: add all cpu internals
	regX        byte
	regY        byte
	accumulator byte
	statusReg   StatusRegister

	cycles int64 // elapsed cpu cycles

	instruction_map [256]func()
}

// Construct a new Cpu
func NewCPU() CPU {
	return CPU{}
}

func (c *CPU) map_instructions() {
	c.instruction_map = [256]func(){0x00: c.Instruction_AAA}
}

// do one step
func (c *CPU) Step() {}

// run indefinitly
func (c *CPU) Run() {}

func (c *CPU) Reset() {
	//TODO: Set registers and internal state
	//TODO: initialize memory
}

func (c *CPU) Powerup() {
	//TODO: Set registers and internal state
	//TODO: initialize memory
}

func (c *CPU) ToString() string {
	return fmt.Sprintf("CPU: %d", c.regX)
}

// implementation of instructions
func (c *CPU) Instruction_AAA() {}

func (c *CPU) aaa() {
	fmt.Println("Instruction aaa called")
}
