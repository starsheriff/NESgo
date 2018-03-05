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

// Status Register of the CPU
// 7 6 5 4 3 2 1 0
// | | | | | | | |
// N V   B D I Z C
//
// C - CarryFlag
// Z - ZeroFlag
// I - InterruptDisable
// D - Decimal Mode
// B - BreakCommand
// V - Overflow Flag
// N - Negative Flag
type StatusRegister byte

//TODO: proper namespacing and type safety?
const (
	CarryFlag        uint = 0
	ZeroFlag         uint = 1
	InterruptDisable uint = 2
	DecimalMode      uint = 3
	BreakCommand     uint = 4
	UnusedFlag       uint = 5
	OverflowFlag     uint = 6
	NegativeFlag     uint = 7
)

// Set sets the whole status register to the given `val` in one operation.
func (c *StatusRegister) Set(val byte) {
	*c = StatusRegister(val)
}

func (c *StatusRegister) setBit(n uint) {
	*c = *c | StatusRegister(1<<n)
}

func (c *StatusRegister) unsetBit(n uint) {
	*c = *c &^ StatusRegister(1<<n)
}

func (c *StatusRegister) SetCarryFlag() {
	c.setBit(CarryFlag)
}

func (c *StatusRegister) UnsetCarryFlag() {
	//TODO: fixme
	c.unsetBit(CarryFlag)
}

func (c *StatusRegister) SetZeroFlag() {
	c.setBit(ZeroFlag)
}

func (c *StatusRegister) UnsetZeroFlag() {
	c.unsetBit(ZeroFlag)
}

//TODO: complete setters and unsetters... really? No generics? That sucks.

// CPU represents the 6502 CPU of the NES system
type CPU struct {
	//TODO: add all cpu internals
	regX        byte
	regY        byte
	accumulator byte
	statusReg   StatusRegister

	cycles int64 // elapsed cpu cycles

	instructionMap [256]func()
}

// NewCPU construct a new CPU object. It is designed in a way that resembles
// a physical silicon chip as close as possible from a users point of view.
// This means the CPU is in a state identical to the silicon version as
// explained in the chip's datasheet.
//
// In particular, it means that the user is responsible in calling the powerup
// and other required initialisation steps as well as connect the CPU to other
// peripherials.
func NewCPU() CPU {
	cpu := CPU{}
	cpu.mapInstructions()
	return cpu
}

// Create the
func (c *CPU) mapInstructions() {
	c.instructionMap = [256]func(){
		//TODO:complete
		0x00: c.aaa,
	}
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

func (c *CPU) aaa() {
	fmt.Println("Instruction aaa called")
}
