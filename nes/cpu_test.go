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
	// simplest case
	var r StatusRegister = 0x00

	r.SetCarryFlag()
	if r != 0x01 {
		t.Errorf("StatusRegister: expected 0x01, got %q", r)
	}

	r.UnsetCarryFlag()
	if r != 0x00 {
		t.Errorf("StatusRegister: expected 0x00, got %q", r)
	}

	// Test that other bits are not set or cleared
	r = 0x04

	r.SetCarryFlag()
	if r != 0x05 {
		t.Errorf("StatusRegister: expected 0x05, got %q", r)
	}

	r.UnsetCarryFlag()
	if r != 0x04 {
		t.Errorf("StatusRegister: expected 0x04, got %q", r)
	}
}
