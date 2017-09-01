package assembler

import "strings"

type cInstruction struct {
	src  string
	dest string
	comp string
	jump string
}

func (c cInstruction) Source() string { return c.src }
func (c cInstruction) String() string {
	output := "111"

	return output
}

func NewCInstruction(src string) instruction {
	c := cInstruction{src, "", "", ""}

	jumpSplit := strings.Split(src, ";")
	if len(jumpSplit) > 1 {
		c.jump = jumpSplit[1]
	}
	remainder := jumpSplit[0]

	destSplit := strings.Split(remainder, "=")
	c.comp = destSplit[0]
	if len(destSplit) > 1 {
		c.dest = destSplit[0]
		c.comp = destSplit[1]
	}

	return c
}
