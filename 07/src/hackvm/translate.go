package hackvm

import (
	"fmt"
	"log"
	"strings"
)

func Translate(vmCode string) (asmCode string) {
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
