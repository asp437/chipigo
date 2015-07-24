package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"time"
)

type CHIP8Console_i interface {
	init(str string)
	loop()
	tick()
}

type CHIP8Console struct {
	mem    CHIP8Memory_i
	cpu    CHIP8CPU_i
	gpu    CHIP8GPU_i
	input  CHIP8Input_i
	sound  CHIP8Sound_i
	window *glfw.Window
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
	console.window = console.gpu.init()
	console.input.init(console.window)
	console.sound.init()
}

func (console *CHIP8Console) loop() {
	cycles_per_second := 120.0
	one_cycle_in_ms := 1.0 / cycles_per_second
	last_time := glfw.GetTime()
	new_time := glfw.GetTime()
	unprocessed := 0.0
	for {
		last_time = new_time
		new_time = glfw.GetTime()
		unprocessed += new_time - last_time
		glfw.PollEvents()
		if console.window.GetKey(glfw.KeyEscape) == glfw.Press || console.window.ShouldClose() {
			break
		}
		for unprocessed > one_cycle_in_ms {
			glfw.PollEvents()
			if console.window.GetKey(glfw.KeyEscape) == glfw.Press || console.window.ShouldClose() {
				break
			}
			console.tick()
			unprocessed -= one_cycle_in_ms
		}
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
