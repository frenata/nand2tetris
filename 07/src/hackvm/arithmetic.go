package hackvm

import (
	"errors"
	"fmt"
)

type operation struct {
	src string
}

func (o operation) Source() string { return o.src }

func NewOperation(line string) (vmInstruction, error) {
	switch line {
	case "add":
		return add{operation{line}}, nil
	default:
		return nil, errors.New("unrecognized operation: " + line)
	}
	//return operation{line}, nil
}

type add struct {
	operation
}

func (o add) Output() string {
	return fmt.Sprintf(
		"// %s\n"+
			"@SP\nA=M\n"+
			"A=A-1\nD=M\n"+
			"A=A-1\nM=M+D\n"+
			"A=A+1\nD=A\n"+
			"@SP\nM=D",
		o.src)
}
