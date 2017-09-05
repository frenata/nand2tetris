package hackvm

type constant struct {
	n int
}

func (c constant) Push() string{
	return "@" + n + "\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n" +
		"@SP\n" +
		"A=A+1\n"
}
