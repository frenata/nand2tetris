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
		return nil, errors.New("unrecognized operation")
	}
	//return operation{line}, nil
}

type add struct {
	operation
}

func (o add) Output() string {
	return fmt.Sprintf(
		"// %s\n@SP\nA=A-1\nD=M\nA=A-1\nM=M+D\nA=A+1",
		o.src)
}
