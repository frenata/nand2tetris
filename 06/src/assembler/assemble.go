package assembler

import (
	"fmt"
	"strings"
)

// Assemble takes as input a string representing a program written in Hack
// and outputs a string representing the machine code for the Hack computer.
func Assemble(program string) string {
	lines := strings.Split(program, "\n")
	output := ""

	for _, line := range lines {
		ins := parse(line)
		if ins == nil {
			continue
		}
		output = fmt.Sprintf("%s%s\n", output, ins.String())
	}

	return output
}
