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
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@$RIP$1
D=A
@$EQ
0;JMP
($RIP$1)
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@$RIP$2
D=A
@$EQ
0;JMP
($RIP$2)
// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@$RIP$3
D=A
@$EQ
0;JMP
($RIP$3)
// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@$RIP$4
D=A
@$LT
0;JMP
($RIP$4)
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@$RIP$5
D=A
@$LT
0;JMP
($RIP$5)
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// lt
@$RIP$6
D=A
@$LT
0;JMP
($RIP$6)
// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@$RIP$7
D=A
@$GT
0;JMP
($RIP$7)
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@$RIP$8
D=A
@$GT
0;JMP
($RIP$8)
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// gt
@$RIP$9
D=A
@$GT
0;JMP
($RIP$9)
// push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 53
@53
D=A
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

// push constant 112
@112
D=A
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

// neg
@SP
A=M-1
M=-M

// and
@SP
AM=M-1
D=M
A=A-1
M=M&D

// push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
// or
@SP
AM=M-1
D=M
A=A-1
M=M|D

// bitwise not
@SP
A=M-1
M=!M

