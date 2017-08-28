// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

@color
M=0

@8192
D=A
@rows
M=D

(WAIT)
@KBD
D=M
@BLACK
D;JNE
@WHITE
0;JMP

(BLACK)
@color
D=M
@WAIT
D+1;JEQ // don't paint if color was already -1
@color
M=-1
@PAINT
0;JMP

(WHITE)
@color
D=M
@WAIT
D;JEQ // don't paint if color was already 0
@color
M=0
@PAINT
0;JMP

(PAINT)
// count up rows from 0 and start row addresses from SCREEN root
@i
M=0
@SCREEN
D=A
@current
M=D

(LOOP)
// if rows - i == 0, stop
@rows
D=M
@i
D=D-M
M=M+1
@WAIT
D;JEQ

// paint the current row and increment it
@color
D=M
@current
A=M
M=D
@current
M=M+1

@LOOP
0;JMP
