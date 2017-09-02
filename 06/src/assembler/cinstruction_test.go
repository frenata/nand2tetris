package assembler

import "testing"

func TestJump(t *testing.T) {
	input := "0;JMP"
	expected := "111"
	actual := NewCInstruction(input).(cInstruction).Jump()

	if actual != expected {
		t.Fatalf("%s did not produce %s: instead %s\n", input, expected, actual)
	}
}

func TestDest(t *testing.T) {
	input := "DM=D+M;JGE"
	expected := "011"
	actual := NewCInstruction(input).(cInstruction).Dest()

	if actual != expected {
		t.Fatalf("%s did not produce %s: instead %s\n", input, expected, actual)
	}
}

func TestComp(t *testing.T) {
	input := "DM=D+M;JGE"
	expected := "1000010"
	actual := NewCInstruction(input).(cInstruction).Comp()

	if actual != expected {
		t.Fatalf("%s did not produce %s: instead %s\n", input, expected, actual)
	}
}
