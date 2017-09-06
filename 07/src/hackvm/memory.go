package hackvm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type memoryAccess struct {
	src     string
	push    bool
	segment string
	index   int
}

var segments map[string]string = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
	"temp":     "5",
}

func NewMemoryAccess(line string) (vmInstruction, error) {
	tokens := strings.Split(line, " ")
	if len(tokens) < 3 {
		return nil, errors.New("no index")
	} else if len(tokens) < 2 {
		return nil, errors.New("no segment specified")
	}

	index, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, err
	}

	segment, ok := segments[tokens[1]]
	if !ok {
		return nil, errors.New("segment not recognized")
	}

	push := false
	if tokens[0] == "push" {
		push = true
	}

	return memoryAccess{line, push, segment, index}, err
}

func (m memoryAccess) Source() string { return m.src }
func (m memoryAccess) Output() string {
	if m.push {
		return m.pushOutput()
	} else {
		return m.popOutput()
	}
}

func (m memoryAccess) pushOutput() string {
	return fmt.Sprintf("// %s\n"+
		"@%d\nD=A\n"+ // save index
		"@%s\nA=M+D\nD=M\n"+ // access memory location and save to D
		"@SP\nA=M\nM=D\n"+ // push to stack
		"@SP\nA=A+1\n", // increment stack
		m.src, m.index, m.segment)
}

func (m memoryAccess) popOutput() string {
	return fmt.Sprintf("// %s\n"+
		"@%d\nD=A\n"+ // save index
		"@%s\nD=M+D\n@R13\nM=D\n"+ // save memory address to R13
		"@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n", // pop to R13 address
		m.src, m.index, m.segment)
}
