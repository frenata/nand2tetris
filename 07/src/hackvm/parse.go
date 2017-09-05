package hackvm

import "strings"

type vmInstruction interface {
	Source() string
	Output() string
}

func parse(line string) (vmInstruction, error) {
	// strip all whitespace and comments
	line = strings.Split(line, "/")[0]
	line = strings.TrimSpace(line)

	switch {
	// if constant return constant
	case strings.Contains(line, "constant"):
		return NewConstant(line)
		// if single word return operation
	case len(line) > 0 && !strings.Contains(line, " "):
		return NewOperation(line)
	// if push/pop return virtual memory access
	// else whitespace or comment return nil
	default:
		return nil, nil
	}
}
