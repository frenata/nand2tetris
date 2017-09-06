package hackvm

import (
	"fmt"
	"log"
	"strings"
)

func Translate(vmCode string) string {
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
