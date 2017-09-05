package hackassembly

import (
	"fmt"
	"testing"
)

func TestParseAInstruction(t *testing.T) {
	input := "@45"
	expected := fmt.Sprintf("%016b", 45)
	actual := parse(input).String()

	if actual != expected {
		t.Fatalf("%s failed to parse as %s: instead %s\n", input, expected, actual)
	}
}

func TestParseCInstruction(t *testing.T) {
	input := "D=D+M"
	expected := "1111000010010000"
	actual := parse(input).String()

	if actual != expected {
		t.Fatalf("%s failed to parse as %s: instead %s\n", input, expected, actual)
	}
}

func TestWhitespace(t *testing.T) {
	input := "	   // test comment"
	actual := parse(input)

	if actual != nil {
		t.Fatalf("%s failed to parse: instead %s\n", input, actual)
	}
}

func TestLabels(t *testing.T) {
	input := "(LOOP) // comment"
	actual := parse(input).String()
	expected := "LOOP"

	if actual != expected {
		t.Fatalf("%s failed to parse as %s: instead %s\n", input, expected, actual)
	}
}
