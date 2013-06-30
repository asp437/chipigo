package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Printf("You must send ROM name. Example\n chipigo maze.rom\n")
	}
	if flag.Arg(0) == "-d" { // Make disasm of rom
		rom_name := flag.Arg(1)
		disasm_rom(rom_name)
	} else {
		_, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		console := CHIP8Console_i(new(CHIP8Console))
		console.init(flag.Arg(0))
		console.loop()
	}
}
