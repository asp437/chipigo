package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("You must send ROM name. Example\n chipigo maze.rom\n")
	}
	_, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	console := CHIP8Console_i(new(CHIP8Console))
	console.init(flag.Arg(0))
	console.loop()
}
