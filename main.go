package main

func main() {
	console := CHIP8Console_i(new(CHIP8Console))
	console.init("test.rom")
	console.loop()
}
