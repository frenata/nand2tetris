package main

import (
	"hackvm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("No file specified")
	}
	filename := args[1]
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	output := hackvm.Translate(string(file), name)
	binary, err := os.Create(name + ".asm")
	if err != nil {
		panic(err)
	}
	defer binary.Close()

	n, err := binary.WriteString(output)
	if err != nil {
		log.Printf("%d bytes written", n)
		panic(err)
	}
}
