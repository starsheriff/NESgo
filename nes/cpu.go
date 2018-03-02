package nes

//      - cycles
//
//

const ClockDivider byte = 12


type instruction struct {
    name string
    size byte
    addressing_mode byte
}

// table of instructions, indexed by optcode
var instructions = [256]instruction {
    //TODO: complete with full instruction set
    0x00 : {"AAA", 3, 3},
    0x01 : {"ABC", 1, 1},
}


type Cpu struct{
    //TODO: add all cpu internals
    reg_x byte
}

// do one step
func (c *Cpu) Step() {}
// run indefinitly
func (c *Cpu) Run() {}

func (c *Cpu) Reset() {}
func (c *Cpu) Powerup() {}

// implementation of instructions
func (c *Cpu) Instruction_AAA() {}

