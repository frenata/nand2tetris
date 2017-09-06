@START
0;JMP
($EQ)
@R15
M=D
@SP
AM=M-1
D=M
A=A-1
D=D-M
@$EQ_END
D;JEQ
D=-1
($EQ_END)
@SP
A=M-1
M=!D
@R15
A=M
0;JMP
($LT)
@R15
M=D
@SP
AM=M-1
D=M
A=A-1
D=M-D
@$LT_TRUE
D;JLT
D=0
@$LT_END
0;JMP
($LT_TRUE)
D=-1
($LT_END)
@SP
A=M-1
M=D
@R15
A=M
0;JMP
($GT)
@R15
M=D
@SP
AM=M-1
D=M
A=A-1
D=M-D
@$GT_TRUE
D;JGT
D=0
@$GT_END
0;JMP
($GT_TRUE)
D=-1
($GT_END)
@SP
A=M-1
M=D
@R15
A=M
0;JMP
(START)
// push constant 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop static 8
@SP
AM=M-1
D=M
@StaticTest.8
M=D

// pop static 3
@SP
AM=M-1
D=M
@StaticTest.3
M=D

// pop static 1
@SP
AM=M-1
D=M
@StaticTest.1
M=D

// push static 3
@StaticTest.3
D=M
@SP
A=M
M=D
@SP
M=M+1
// push static 1
@StaticTest.1
D=M
@SP
A=M
M=D
@SP
M=M+1
// sub
@SP
AM=M-1
D=M
A=A-1
M=M-D

// push static 8
@StaticTest.8
D=M
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
AM=M-1
D=M
A=A-1
M=M+D

