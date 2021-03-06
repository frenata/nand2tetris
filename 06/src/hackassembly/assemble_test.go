package hackassembly_test

import (
	"hackassembly"
	"io/ioutil"
	"testing"
)

func TestAssemble(t *testing.T) {
	var testFiles map[string]string = map[string]string{
		"../../tests/add/Add.asm":    "../../tests/add/Add.cmp",
		"../../tests/max/MaxL.asm":   "../../tests/max/MaxL.cmp",
		"../../tests/rect/RectL.asm": "../../tests/rect/RectL.cmp",
		"../../tests/pong/PongL.asm": "../../tests/pong/PongL.cmp",
		"../../tests/max/Max.asm":    "../../tests/max/Max.cmp",
		"../../tests/rect/Rect.asm":  "../../tests/rect/Rect.cmp",
		"../../tests/pong/Pong.asm":  "../../tests/pong/Pong.cmp",
	}

	for asmFile, cmpFile := range testFiles {
		input, err := ioutil.ReadFile(asmFile)
		if err != nil {
			panic(err)
		}
		expected, err := ioutil.ReadFile(cmpFile)
		if err != nil {
			panic(err)
		}
		actual := hackassembly.Assemble(string(input))

		if actual != string(expected) {
			t.Logf("Actual:\n%s", actual)
			t.Logf("Expected:\n%s", expected)
			t.Fatalf("Output of Assembler did not match the contents of %s", cmpFile)
		}
	}
}
