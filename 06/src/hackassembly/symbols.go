package hackassembly

import (
	"fmt"
	"strconv"
	"strings"
)

type label struct {
	src  string
	name string
}

func NewLabel(src string) instruction {
	return label{src, src[1 : len(src)-1]}
}
func (a label) Source() string { return a.src }
func (a label) String() string { return a.name }

// Stages a symbolic assembly program, by stripping all symbols into raw numbers.
func Stage(program string) string {
	table := stageLabels(program)
	nextFreeAddress := 16

	lines := strings.Split(program, "\n")
	output := ""

	for _, line := range lines {
		ins := parse(line)

		switch ins.(type) {
		case aInstruction:
			target := ins.Source()[1:]
			if _, err := strconv.Atoi(target); err == nil {
				output = fmt.Sprintf("%s%s\n", output, ins.Source())
				continue
			}

			address, ok := table[target]
			if !ok {
				address = strconv.Itoa(nextFreeAddress)
				table[target] = address
				nextFreeAddress++
			}
			output = fmt.Sprintf("%s@%s\n", output, address)

		case cInstruction:
			output = fmt.Sprintf("%s%s\n", output, ins.Source())
		}
	}

	return output
}

func stageLabels(program string) map[string]string {
	table := blankTable()

	lines := strings.Split(program, "\n")
	lineNum := -1

	for _, line := range lines {
		ins := parse(line)
		switch ins.(type) {
		case aInstruction, cInstruction:
			lineNum++
		case label:
			table[ins.String()] = strconv.Itoa(lineNum + 1)
		}
	}

	return table
}

func blankTable() map[string]string {
	return map[string]string{
		"R0": "0", "R1": "1", "R2": "2", "R3": "3",
		"R4": "4", "R5": "5", "R6": "6", "R7": "7",
		"R8": "8", "R9": "9", "R10": "10", "R11": "11",
		"R12": "12", "R13": "13", "R14": "14", "R15": "15",

		"SCREEN": "16384", "KBD": "24576",

		"SP": "0", "LCL": "1", "ARG": "2", "THIS": "3", "THAT": "4",
	}
}
