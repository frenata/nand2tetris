package hackvm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type constant struct {
	src string
	n   int
}

func NewConstant(line string) (vmInstruction, error) {
	tokens := strings.Split(line, " ")
	if len(tokens) < 3 {
		return nil, errors.New("no index")
	}
	n, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, err
	}

	return constant{line, n}, err
}

func (c constant) Source() string { return c.src }
func (c constant) Output() string {
	return fmt.Sprintf(
		"// %s\n@%d\nD=A\n@SP\nM=D\n@SP\nA=A+1",
		c.src, c.n)
}
