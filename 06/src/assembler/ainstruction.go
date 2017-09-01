package assembler

import (
	"fmt"
	"strconv"
)

type aInstruction struct {
	src     string
	address int
}

func NewAInstruction(src string) instruction {
	n, err := strconv.Atoi(src[1:])
	if err != nil {
		n = -1
	}
	return aInstruction{src, n}
}
func (a aInstruction) Source() string { return a.src }
func (a aInstruction) String() string {
	return fmt.Sprintf("%016b", a.address)
}
