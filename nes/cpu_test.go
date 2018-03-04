package nes

import "testing"
import "fmt"

func TestInstructionsTable(t *testing.T) {
	expected := instruction{"ABC", 1, 1}
	result := instructions[1]

	fmt.Println(instructions[5], result)

	if result != expected {
		t.Errorf("Instruction is wrong")
	}
}

func TestCPUToString(t *testing.T) {
	c := NewCPU()
	res := c.ToString()

	fmt.Println(res)
}

func TestStatusRegister(t *testing.T) {
	var r StatusRegister = 0x00

	r.SetCarryFlag()

	if r != 0x01 {
		t.Errorf("StatusRegister: expected 0x01, got %q", r)
	}
}
