// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

@R0
D=M
@n
M=D

@R1
D=M
@m
M=D

@total
M=0

(LOOP)
//check if m is <= 0
@m
D=M
@OUTPUT
D;JLE

@n
D=M
@total
M=M+D
@m
M=M-1

@LOOP
0;JMP

(OUTPUT)
@total
D=M
@R2
M=D

(END)
@END
0;JMP
