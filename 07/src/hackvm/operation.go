package hackvm

import (
	"errors"
	"fmt"
)

type operation struct {
	src string
}

var operations map[string]func() string = map[string]func() string{
	"add": add,
	"sub": sub,
	"and": and,
	"or":  or,
	"neg": neg,
	"not": not,
	"eq":  eq,
	"lt":  lt,
	"gt":  gt,
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

// binary operations
func add() string {
	return "// add\n" +
		"@SP\nAM=M-1\n" +
		"D=M\nA=A-1\nM=M+D\n"
}

func sub() string {
	return "// sub\n" +
		"@SP\nAM=M-1\n" +
		"D=M\nA=A-1\nM=M-D\n"
}

func and() string {
	return "// and\n" +
		"@SP\nAM=M-1\n" +
		"D=M\nA=A-1\nM=M&D\n"
}

func or() string {
	return "// or\n" +
		"@SP\nAM=M-1\n" +
		"D=M\nA=A-1\nM=M|D\n"
}

// unary operations
func neg() string {
	return "// neg\n" +
		"@SP\nA=M-1\n" +
		"M=-M\n"
}

func not() string {
	return "// bitwise not\n" +
		"@SP\nA=M-1\n" +
		"M=!M\n"
}

// comparison operations
func eq() string {
	returnAddr := fmt.Sprintf("$RIP$%d", jumpAddr)
	return "// eq\n" +
		"@" + returnAddr + "\n" +
		"D=A\n@$EQ\n0;JMP\n" +
		"(" + returnAddr + ")"
}

func lt() string {
	returnAddr := fmt.Sprintf("$RIP$%d", jumpAddr)
	return "// lt\n" +
		"@" + returnAddr + "\n" +
		"D=A\n@$LT\n0;JMP\n" +
		"(" + returnAddr + ")"
}

func gt() string {
	returnAddr := fmt.Sprintf("$RIP$%d", jumpAddr)
	return "// gt\n" +
		"@" + returnAddr + "\n" +
		"D=A\n@$GT\n0;JMP\n" +
		"(" + returnAddr + ")"
}
