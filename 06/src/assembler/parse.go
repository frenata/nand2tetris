package assembler

import "strings"

type instruction interface {
	Source() string
	String() string
}

func parse(line string) instruction {
	nocomments := strings.Split(line, "/")[0]
	ins := strings.TrimSpace(nocomments)

	switch {
	case ins == "": // just whitespace
		return nil
	case strings.HasPrefix(ins, "("): // label
		return nil // TODO
	case strings.HasPrefix(ins, "@"):
		return NewAInstruction(ins)
	default: // C instruction
		return NewCInstruction(ins)
	}
}
