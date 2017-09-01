package assembler_test

import (
	"assembler"
	"io/ioutil"
	"testing"
)

func TestAssemble(t *testing.T) {
	asmFile := "../../tests/add/Add.asm"
	cmpFile := "../../tests/add/Add.cmp"

	input, err := ioutil.ReadFile(asmFile)
	if err != nil {
		panic(err)
	}
	expected, err := ioutil.ReadFile(cmpFile)
	if err != nil {
		panic(err)
	}
	actual := assembler.Assemble(string(input))

	if actual != string(expected) {
		t.Logf("Actual:\n%s", actual)
		t.Logf("Expected:\n%s", expected)
		t.Fatalf("Output of Assembler did not match the contents of %s", cmpFile)
	}
}
