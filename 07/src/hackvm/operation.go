package hackvm

import (
	"errors"
	"strconv"
)

type operation struct {
	src string
}

var operations map[string]func() string = map[string]func() string{
	"add": add,
	"eq":  eq,
}

var jumpAddr int

func (o operation) Source() string { return o.src }
func (o operation) Output() string { return operations[o.src]() }

func NewOperation(line string) (vmInstruction, error) {
	jumpAddr++
	if _, ok := operations[line]; !ok {
		return nil, errors.New("unrecognized operation: " + line)
	}
	return operation{line}, nil
	//return operation{line}, nil
}

func init() {
	jumpAddr = 0
}

func add() string {
	return "// add\n" +
		"@SP\nA=M\n" +
		"A=A-1\nD=M\n" +
		"A=A-1\nM=M+D\n" +
		"A=A+1\nD=A\n@SP\nM=D" // increment SP
}

func eq() string {
	return "// eq\n" +
		"@SP\nA=M\nA=A-1\nD=M\n" + // pop to D
		"A=A-1\nD=D-M\n" + // D is zero if equal
		"@eq-" + strconv.Itoa(jumpAddr) + "\nD;JEQ\n" +
		"@SP\nA=M\nA=A-1\nA=A-1\nM=0\n" +
		"@eq-" + strconv.Itoa(jumpAddr) + "-end\n0;JMP\n" +
		"(eq-" + strconv.Itoa(jumpAddr) + ")\n" +
		"@SP\nA=M\nA=A-1\nA=A-1\nM=-1\n" +
		"@eq-" + strconv.Itoa(jumpAddr) + "-end\n0;JMP\n" +
		"(eq-" + strconv.Itoa(jumpAddr) + "-end)\n" +
		"@SP\nM=M-1\n"
}
