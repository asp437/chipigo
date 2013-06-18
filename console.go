package main

import (
	"github.com/go-gl/glfw"
	"time"
)

type CHIP8Console_i interface {
	init(str string)
	loop()
	tick()
}

type CHIP8Console struct {
	mem   CHIP8Memory_i
	cpu   CHIP8CPU_i
	gpu   CHIP8GPU_i
	input CHIP8Input_i
	sound CHIP8Sound_i
}

func (console *CHIP8Console) init(str string) {
	console.cpu = new(CHIP8CPU)
	console.mem = new(CHIP8Memory)
	console.gpu = new(CHIP8GPU)
	console.input = new(CHIP8Input)
	console.sound = new(CHIP8Sound)

	console.cpu.init()
	console.mem.init()
	console.mem.read_rom(str)
	console.gpu.init()
	console.input.init()
	console.sound.init()
}

func (console *CHIP8Console) loop() {
	for {
		glfw.PollEvents()
		if glfw.Key(glfw.KeyEsc) == glfw.KeyPress || glfw.WindowParam(glfw.Opened) == 0 {
			break
		}
		console.tick()
		time.Sleep(10 * time.Millisecond)
	}
	glfw.Terminate()
}

func (console *CHIP8Console) tick() {
	console.input.tick()
	console.cpu.tick(console)
	console.gpu.render()
	console.sound.tick()
}
