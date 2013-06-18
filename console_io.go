package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
)

type CHIP8GPU_i interface {
	clear_screen()
	render()
	init()
	draw_line8(x, y int8, line uint8) Registr // Return new value of VF
}

type CHIP8GPU struct {
	pic   [][]uint8 // Screen can display only black and some 1 color
	w, h  int       // width and height of screen
	color uint32    // color for filled pixels
} // Not full implemented yet

// The interpreter reads n bytes from memory, starting at the address stored in I.
// These bytes are then displayed as sprites on screen at coordinates (Vx, Vy).
// Sprites are XORed onto the existing screen.
// If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0.
// If the sprite is positioned so part of it is outside the coordinates of the display,
// it wraps around to the opposite side of the screen.
// See instruction 8XY3 for more information on XOR,
// and section 2.4, Display, for more information on the Chip-8 screen and sprites.
func (gpu *CHIP8GPU) draw_line8(x, y int8, line uint8) Registr {
	ret := 0
	if y >= 32 {
		y -= 32
	}
	for y < 0 {
		y += 32
	}
	for i := 0; i < 8; i++ {
		xp := x + int8(i)
		if xp >= 64 {
			xp -= 64
		}
		for xp < 0 {
			xp += 64
		}
		new_pix := (gpu.pic[xp][y] & 1) ^ ((line >> uint(7-i)) & 1)
		if gpu.pic[xp][y] == 1 && new_pix == 0 {
			ret = 1
		}
		gpu.pic[xp][y] = new_pix
	}
	return Registr(ret)
}

func (gpu *CHIP8GPU) clear_screen() {
	for x := 0; x < gpu.w; x++ {
		for y := 0; y < gpu.h; y++ {
			gpu.pic[x][y] = 0
		}
	}
}

func (gpu *CHIP8GPU) render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	for x := 0; x < gpu.w; x++ {
		for y := 0; y < gpu.h; y++ {
			if gpu.pic[x][y] != 0 {
				gl.Begin(gl.QUADS)
				gl.Color3ub(uint8(gpu.color&0xFF0000>>16), uint8(gpu.color&0x00FF00>>8), uint8(gpu.color&0x0000FF))
				gl.Vertex2i(x, y)
				gl.Vertex2i(x+1, y)
				gl.Vertex2i(x+1, y+1)
				gl.Vertex2i(x, y+1)
				gl.End()
			}
		}
	}
	glfw.SwapBuffers()
}

func (gpu *CHIP8GPU) init() {
	gpu.w = 64
	gpu.h = 32
	gpu.pic = make([][]uint8, gpu.w)
	for x := 0; x < gpu.w; x++ {
		gpu.pic[x] = make([]uint8, gpu.h)
		for y := 0; y < gpu.h; y++ {
			gpu.pic[x][y] = 0
		}
	}
	gpu.color = 0x11FF11 // Some kind of green

	glfw.Init()
	glfw.OpenWindow(640, 320, 0, 0, 0, 0, 0, 0, glfw.Windowed)
	gl.Init()

	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
	gl.ClearDepth(1.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, 64, 32, 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

type CHIP8Sound_i interface {
	init()
	turn_beep(val bool) // Change beep status from true to false and from false to true
	tick()
}

type CHIP8Sound struct {
	turn_on bool
} // Not implemented

func (sound *CHIP8Sound) init() {
	sound.turn_on = false
}

func (sound *CHIP8Sound) turn_beep(val bool) {
	sound.turn_on = val
}

func (sound *CHIP8Sound) tick() {}

type CHIP8Input_i interface {
	init()
	is_pressed(key uint8) bool
	tick()
}

type CHIP8Input struct {
	keys int
}

func (input *CHIP8Input) init() {
	input.keys = -255
}

func (input *CHIP8Input) is_pressed(key uint8) bool {
	return (input.keys == int(key))
}

func (input *CHIP8Input) tick() {
	input.keys = -255
	if glfw.Key('1') == glfw.KeyPress {
		input.keys = 0x1
	} else if glfw.Key('2') == glfw.KeyPress {
		input.keys = 0x2
	} else if glfw.Key('3') == glfw.KeyPress {
		input.keys = 0x3
	} else if glfw.Key('4') == glfw.KeyPress {
		input.keys = 0xC
	} else if glfw.Key('Q') == glfw.KeyPress {
		input.keys = 0x4
	} else if glfw.Key('W') == glfw.KeyPress {
		input.keys = 0x5
	} else if glfw.Key('E') == glfw.KeyPress {
		input.keys = 0x6
	} else if glfw.Key('R') == glfw.KeyPress {
		input.keys = 0xD
	} else if glfw.Key('A') == glfw.KeyPress {
		input.keys = 0x7
	} else if glfw.Key('S') == glfw.KeyPress {
		input.keys = 0x8
	} else if glfw.Key('D') == glfw.KeyPress {
		input.keys = 0x9
	} else if glfw.Key('F') == glfw.KeyPress {
		input.keys = 0xE
	} else if glfw.Key('Z') == glfw.KeyPress {
		input.keys = 0xA
	} else if glfw.Key('X') == glfw.KeyPress {
		input.keys = 0x0
	} else if glfw.Key('C') == glfw.KeyPress {
		input.keys = 0xB
	} else if glfw.Key('V') == glfw.KeyPress {
		input.keys = 0xF
	}
}
