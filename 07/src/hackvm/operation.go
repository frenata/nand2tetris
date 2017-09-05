package hackvm

import "errors"

type operation struct {
	src string
}

func (o operation) Source() string { return o.src }
func (o operation) Output() string { return operations[o.src] }

func NewOperation(line string) (vmInstruction, error) {
	if _, ok := operations[line]; !ok {
		return nil, errors.New("unrecognized operation: " + line)
	}
	return operation{line}, nil
	//return operation{line}, nil
}

var operations map[string]string = map[string]string{
	"add": add,
}

const add string = "// add\n" +
	"@SP\nA=M\n" +
	"A=A-1\nD=M\n" +
	"A=A-1\nM=M+D\n" +
	"A=A+1\nD=A\n" +
	"@SP\nM=D"
