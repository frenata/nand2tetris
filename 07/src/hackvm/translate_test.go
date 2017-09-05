package hackvm_test

import (
	"hackvm"
	"io/ioutil"
	"testing"
)

func TestSimpleTranslate(t *testing.T) {
	t.Log(hackvm.Translate("eq"))
}

func TestTranslate(t *testing.T) {
	var testFiles map[string]string = map[string]string{
		"../../tests/StackArithmetic/SimpleAdd/SimpleAdd.vm": "../../tests/StackArithmetic/SimpleAdd/SimpleAdd.asm",
	}

	for vmFile, asmFile := range testFiles {
		input, err := ioutil.ReadFile(vmFile)
		if err != nil {
			panic(err)
		}
		expected, err := ioutil.ReadFile(asmFile)
		if err != nil {
			panic(err)
		}
		actual := hackvm.Translate(string(input))

		if actual != string(expected) {
			t.Logf("Actual:\n%s", actual)
			t.Logf("Expected:\n%s", expected)
			t.Fatalf("Output of VM translator did not match the contents of %s", asmFile)
		}
	}
}
