package hackvm

import (
	"fmt"
	"log"
	"strings"
)

var name string = "NO_NAME"

func Translate(vmCode string, fileName string) string {
	name = fileName
	asmCode := bootstrap()
	lines := strings.Split(vmCode, "\n")

	for i, line := range lines {
		ins, err := parse(line)
		if err != nil {
			log.Fatalf("Syntax error in line %d:\n%s\n%s\n", i+1, line, err)
		} else if ins == nil { // ignore comments/whitespace
			continue
		}

		asmCode = fmt.Sprintf("%s%s\n", asmCode, ins.Output())
	}

	return asmCode
}

// define some common assembly functions and start the program
func bootstrap() string {
	return "@START\n0;JMP\n" +
		eqFunction +
		ltFunction +
		gtFunction +
		"(START)\n"
}

// pops top 2 args and checks equality
// pushes true: -1 or false: 0 to stack
// returns to addr stored in D
const eqFunction string = "($EQ)\n" +
	"@R15\nM=D\n" + // save return addr
	"@SP\nAM=M-1\nD=M\nA=A-1\nD=D-M\n" + // calculate arg2-arg1
	"@$EQ_END\nD;JEQ\n" +
	"D=-1\n" +
	"($EQ_END)\n@SP\nA=M-1\nM=!D\n" + // push to stack
	"@R15\nA=M\n0;JMP\n" // return

// pops top 2 args and checks less than
// pushes true: -1 or false: 0 to stack
// returns to addr stored in D
const ltFunction string = "($LT)\n" +
	"@R15\nM=D\n" + // save return addr
	"@SP\nAM=M-1\nD=M\nA=A-1\nD=M-D\n" + // calculate arg2-arg1
	"@$LT_TRUE\nD;JLT\n" +
	"D=0\n@$LT_END\n0;JMP\n" +
	"($LT_TRUE)\nD=-1\n" +
	"($LT_END)\n@SP\nA=M-1\nM=D\n" + // push to stack
	"@R15\nA=M\n0;JMP\n" // return

// pops top 2 args and checks greater than
// pushes true: -1 or false: 0 to stack
// returns to addr stored in D
const gtFunction string = "($GT)\n" +
	"@R15\nM=D\n" + // save return addr
	"@SP\nAM=M-1\nD=M\nA=A-1\nD=M-D\n" + // calculate arg2-arg1
	"@$GT_TRUE\nD;JGT\n" +
	"D=0\n@$GT_END\n0;JMP\n" +
	"($GT_TRUE)\nD=-1\n" +
	"($GT_END)\n@SP\nA=M-1\nM=D\n" + // push to stack
	"@R15\nA=M\n0;JMP\n" // return
