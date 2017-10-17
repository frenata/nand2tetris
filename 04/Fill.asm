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

(INIT)		// max screen position is keyboard register - 1 
    @KBD
    D=A
    @max
    M=D-1

(RESET)		// offset starts (and resets) to the screen base
    @SCREEN
    D=A
    @offset
    M=D

(LOOP)		
    @KBD	// check keyboard input
    D=M
    @PAINT	// if no key is pressed: D is already white
    D;JEQ
    D=-1	// if key is pressed, set D to black
    
(PAINT)		// paint the current offset
    @offset
    A=M
    M=D
    @offset
    DM=M+1	// increment offset
    
    @max	// reset offset counter if max is reached
    D=D-M
    @RESET
    D;JGE

    @LOOP
    0;JMP
