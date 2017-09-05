package hackassembly_test

import (
	"bytes"
	"hackassembly"
	"io/ioutil"
	"testing"
)

func TestStage(t *testing.T) {
	var testFiles map[string]string = map[string]string{
		"../../tests/max/Max.asm":   "../../tests/max/MaxL.asm",
		"../../tests/rect/Rect.asm": "../../tests/rect/RectL.asm",
		"../../tests/pong/Pong.asm": "../../tests/pong/PongL.asm",
	}

	for symFile, noSymFile := range testFiles {
		input, err := ioutil.ReadFile(symFile)
		if err != nil {
			panic(err)
		}
		expected, err := ioutil.ReadFile(noSymFile)
		if err != nil {
			panic(err)
		}
		expected = bytes.SplitN(expected, []byte{'\n'}, 8)[7]
		actual := hackassembly.Stage(string(input))

		if actual != string(expected) {
			t.Logf("Actual:\n%s", actual)
			t.Logf("Expected:\n%s", expected)
			t.Fatalf("Output of Stage did not match the contents of %s", noSymFile)
		}
	}
}
