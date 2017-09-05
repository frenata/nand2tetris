package hackassembly

import "strings"

type cInstruction struct {
	src  string
	dest string
	comp string
	jump string
}

func (c cInstruction) Source() string { return c.src }
func (c cInstruction) String() string {
	prefix := "111"
	//return prefix + c.Comp() + c.Dest() + c.Jump()
	return prefix + c.Comp() + c.Dest() + c.Jump()
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

func (c cInstruction) Jump() string {
	return map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}[c.jump]
}

func (c cInstruction) Dest() string {
	return map[string]string{
		"":   "000",
		"M":  "001",
		"D":  "010",
		"MD": "011", "DM": "011",
		"A":  "100",
		"AM": "101", "MA": "101",
		"AD": "110", "DA": "110",
		"AMD": "111", "ADM": "111",
		"MAD": "111", "MDA": "111",
		"DMA": "111", "DAM": "111",
	}[c.dest]
}

func (c cInstruction) Comp() string {
	a := "0"
	if strings.Contains(c.comp, "M") {
		a = "1"
	}

	return a + map[string]string{
		"0":   "101010",
		"1":   "111111",
		"-1":  "111010",
		"D":   "001100",
		"A":   "110000",
		"!D":  "001101",
		"!A":  "110001",
		"-D":  "001111",
		"-A":  "110011",
		"D+1": "011111",
		"A+1": "110111",
		"D-1": "001110",
		"A-1": "110010",
		"D+A": "000010", "A+D": "000010",
		"D-A": "010011",
		"A-D": "000111",
		"D&A": "000000", "A&D": "000000",
		"D|A": "010101", "A|D": "010101",
	}[strings.Replace(c.comp, "M", "A", -1)]
}
