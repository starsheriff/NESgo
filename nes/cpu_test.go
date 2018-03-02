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
