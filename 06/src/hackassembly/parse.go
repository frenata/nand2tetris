package hackassembly

import "strings"

type instruction interface {
	Source() string
	String() string
}

func parse(line string) instruction {
	// strip all whitespace and comments
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\t", "", -1)
	line = strings.Split(line, "/")[0]
	line = strings.TrimSpace(line)

	switch {
	case line == "": // just whitespace
		return nil
	case strings.HasPrefix(line, "("): // label
		return NewLabel(line)
	case strings.HasPrefix(line, "@"):
		return NewAInstruction(line)
	default: // C instruction
		return NewCInstruction(line)
	}
}
