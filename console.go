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
	cycles_per_second := 1000.0
	one_cycle_in_ms := 1000.0 / cycles_per_second
	var last_time, new_time, unprocessed float64
	last_time = glfw.Time() * 1000
	new_time = glfw.Time() * 1000
	unprocessed = 0.0
	for {
		last_time = new_time
		new_time = glfw.Time() * 1000
		unprocessed += new_time - last_time
		glfw.PollEvents()
		if glfw.Key(glfw.KeyEsc) == glfw.KeyPress || glfw.WindowParam(glfw.Opened) == 0 {
			break
		}
		for unprocessed > one_cycle_in_ms {
			if glfw.Key(glfw.KeyEsc) == glfw.KeyPress || glfw.WindowParam(glfw.Opened) == 0 {
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
