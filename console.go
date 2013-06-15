package main

type OpCode uint32
type Registr uint8
type CPUTimer uint8

type CHIP8CPU_i interface {
	op_0NNN(op OpCode, console *CHIP8Console) // 0NNN - Calls RCA 1802 program at address NNN. (No implement needed)
	op_00E0(op OpCode, console *CHIP8Console) // 00E0 - Clears the screen.
	op_00EE(op OpCode, console *CHIP8Console) // 00EE - Returns from a subroutine.
	op_1NNN(op OpCode, console *CHIP8Console) // 1NNN - Jumps to address NNN.
	op_2NNN(op OpCode, console *CHIP8Console) // 2NNN - Calls subroutine at NNN.
	op_3XNN(op OpCode, console *CHIP8Console) // 3XNN - Skips the next instruction if VX equals NN.
	op_4XNN(op OpCode, console *CHIP8Console) // 4XNN - Skips the next instruction if VX doesn't equal NN.
	op_5XY0(op OpCode, console *CHIP8Console) // 5XY0 - Skips the next instruction if VX equals VY.
	op_6XNN(op OpCode, console *CHIP8Console) // 6XNN - Sets VX to NN.
	op_7XNN(op OpCode, console *CHIP8Console) // 7XNN - Adds NN to VX.
	op_8XY0(op OpCode, console *CHIP8Console) // 8XY0 - Sets VX to the value of VY.
	op_8XY1(op OpCode, console *CHIP8Console) // 8XY1 - Sets VX to VX or VY.
	op_8XY2(op OpCode, console *CHIP8Console) // 8XY2 - Sets VX to VX and VY.
	op_8XY3(op OpCode, console *CHIP8Console) // 8XY3 - Sets VX to VX xor VY.
	op_8XY4(op OpCode, console *CHIP8Console) // 8XY4 - Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
	op_8XY5(op OpCode, console *CHIP8Console) // 8XY5 - VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	op_8XY6(op OpCode, console *CHIP8Console) // 8XY6 - Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.[2]
	op_8XY7(op OpCode, console *CHIP8Console) // 8XY7 - Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	op_8XYE(op OpCode, console *CHIP8Console) // 8XYE - Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.[2]
	op_9XY0(op OpCode, console *CHIP8Console) // 9XY0 - Skips the next instruction if VX doesn't equal VY.
	op_ANNN(op OpCode, console *CHIP8Console) // ANNN - Sets I to the address NNN.
	op_BNNN(op OpCode, console *CHIP8Console) // BNNN - Jumps to the address NNN plus V0.
	op_CXNN(op OpCode, console *CHIP8Console) // CXNN - Sets VX to a random number and NN.
	op_DXYN(op OpCode, console *CHIP8Console) // DXYN - Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels. Each row of 8 pixels is read as bit-coded (with the most significant bit of each byte displayed on the left) starting from memory location I; I value doesn't change after the execution of this instruction. As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that doesn't happen.
	op_EX9E(op OpCode, console *CHIP8Console) // EX9E - Skips the next instruction if the key stored in VX is pressed.
	op_EXA1(op OpCode, console *CHIP8Console) // EXA1 - Skips the next instruction if the key stored in VX isn't pressed.
	op_FX07(op OpCode, console *CHIP8Console) // FX07 - Sets VX to the value of the delay timer.
	op_FX0A(op OpCode, console *CHIP8Console) // FX0A - A key press is awaited, and then stored in VX.
	op_FX15(op OpCode, console *CHIP8Console) // FX15 - Sets the delay timer to VX.
	op_FX18(op OpCode, console *CHIP8Console) // FX18 - Sets the sound timer to VX.
	op_FX1E(op OpCode, console *CHIP8Console) // FX1E - Adds VX to I.[3]
	op_FX29(op OpCode, console *CHIP8Console) // FX29 - Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
	op_FX33(op OpCode, console *CHIP8Console) // FX33 - Stores the Binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2. (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
	op_FX55(op OpCode, console *CHIP8Console) // FX55 - Stores V0 to VX in memory starting at address I.[4]
	op_FX65(op OpCode, console *CHIP8Console) // FX65 - Fills V0 to VX with values from memory starting at address I.[4]
	init()
}

type CHIP8CPU struct {
	// TODO: Rewrite stack. Place it into console memory
	v  []Registr // V0 - VF registers
	i  uint16    // I register
	pc uint16    // Currently executing address
	sp *Stack    // CPU Stack
	dt CPUTimer  // Delay timer
	st CPUTimer  // Sound timer
}

func (cpu *CHIP8CPU) init() {
	cpu.v = make([]Registr, 16)
	cpu.i = 0x200 // First 0x200 byte are interpreter
	stack := new(Stack)
	stack.init(16)
}

type CHIP8Memory_i interface {
	init()
	read(addr uint32) uint8
	write(addr uint32, val uint8)
	read2(addr uint32) uint16 // Read 2 byte. Special for opcode reading
}

type CHIP8Memory struct {
	data []uint8
}

func (mem *CHIP8Memory) read(addr uint32) uint8 {
	return mem.data[addr]
}

func (mem *CHIP8Memory) read2(addr uint32) uint16 {
	return (uint16(mem.data[addr]) << 8) + uint16(mem.data[addr+1])
}

func (mem *CHIP8Memory) write(addr uint32, val uint8) {
	mem.data[addr] = val
}

func (mem *CHIP8Memory) init() {
	mem.data = make([]uint8, 0x1000)

	mem.data[0x0] = 0xF0 // ****
	mem.data[0x1] = 0x90 // *  *
	mem.data[0x2] = 0x90 // *  *
	mem.data[0x3] = 0x90 // *  *
	mem.data[0x4] = 0xF0 // ****

	mem.data[0x5] = 0x20 //   *
	mem.data[0x6] = 0x60 //  **
	mem.data[0x7] = 0x20 //   *
	mem.data[0x8] = 0x20 //   *
	mem.data[0x9] = 0x70 //  ***

	mem.data[0xA] = 0xF0 // ****
	mem.data[0xB] = 0x10 //    *
	mem.data[0xC] = 0xF0 // ****
	mem.data[0xD] = 0x80 // *
	mem.data[0xE] = 0xF0 // ****

	mem.data[0xF] = 0xF0  // ****
	mem.data[0x10] = 0x10 //    *
	mem.data[0x11] = 0xF0 // ****
	mem.data[0x12] = 0x10 //    *
	mem.data[0x13] = 0xF0 // ****

	mem.data[0x14] = 0x90 // *  *
	mem.data[0x15] = 0x90 // *  *
	mem.data[0x16] = 0xF0 // ****
	mem.data[0x17] = 0x10 //    *
	mem.data[0x18] = 0x10 //    *

	mem.data[0x19] = 0xF0 // ****
	mem.data[0x1A] = 0x80 // *
	mem.data[0x1B] = 0xF0 // ****
	mem.data[0x1C] = 0x10 //    *
	mem.data[0x1D] = 0xF0 // ****

	mem.data[0x1E] = 0xF0 // ****
	mem.data[0x1F] = 0x80 // *
	mem.data[0x20] = 0xF0 // ****
	mem.data[0x21] = 0x90 // *  *
	mem.data[0x22] = 0xF0 // ****

	mem.data[0x23] = 0xF0 // ****
	mem.data[0x24] = 0x10 //    *
	mem.data[0x25] = 0x20 //   *
	mem.data[0x26] = 0x40 //  *
	mem.data[0x27] = 0x40 //  *

	mem.data[0x28] = 0xF0 // ****
	mem.data[0x29] = 0x90 // *  *
	mem.data[0x2A] = 0xF0 // ****
	mem.data[0x2B] = 0x90 // *  *
	mem.data[0x2C] = 0xF0 // ****

	mem.data[0x2D] = 0xF0 // ****
	mem.data[0x2E] = 0x90 // *  *
	mem.data[0x2F] = 0xF0 // ****
	mem.data[0x30] = 0x10 //    *
	mem.data[0x31] = 0xF0 // ****

	mem.data[0x32] = 0xF0 // ****
	mem.data[0x33] = 0x90 // *  *
	mem.data[0x34] = 0xF0 // ****
	mem.data[0x35] = 0x90 // *  *
	mem.data[0x36] = 0x90 // *  *

	mem.data[0x37] = 0xE0 // ***
	mem.data[0x38] = 0x90 // *  *
	mem.data[0x39] = 0xE0 // ***
	mem.data[0x3A] = 0x90 // *  *
	mem.data[0x3B] = 0xE0 // ***

	mem.data[0x3C] = 0xF0 // ****
	mem.data[0x3D] = 0x80 // *
	mem.data[0x3E] = 0x80 // *
	mem.data[0x3F] = 0x80 // *
	mem.data[0x40] = 0xF0 // ****

	mem.data[0x41] = 0xE0 // ***
	mem.data[0x42] = 0x90 // *  *
	mem.data[0x43] = 0x90 // *  *
	mem.data[0x44] = 0x90 // *  *
	mem.data[0x45] = 0xE0 // ***

	mem.data[0x46] = 0xF0 // ****
	mem.data[0x47] = 0x80 // *
	mem.data[0x48] = 0xF0 // ****
	mem.data[0x49] = 0x80 // *
	mem.data[0x4A] = 0xF0 // ****

	mem.data[0x4B] = 0xF0 // ****
	mem.data[0x4C] = 0x80 // *
	mem.data[0x4D] = 0xF0 // ****
	mem.data[0x4E] = 0x80 // *
	mem.data[0x4F] = 0x80 // *
}

type CHIP8GPU_i interface {
	clear_screen()
	render()
	init()
	draw_line8(x, y, line uint8)
}

type CHIP8GPU struct {
	pic   [][]bool // Screen can display only black and some 1 color
	w, h  int      // width and height of screen
	color uint32   // color for filled pixels
} // Not full implemented yet

func (gpu *CHIP8GPU) draw_line8(x, y, line uint8) {

}

func (gpu *CHIP8GPU) clear_screen() {
	for x := 0; x < gpu.w; x++ {
		for y := 0; y < gpu.h; y++ {
			gpu.pic[x][y] = false
		}
	}
}

func (gpu *CHIP8GPU) render() {} // Not implemented

func (gpu *CHIP8GPU) init() {
	gpu.w = 64
	gpu.h = 32
	gpu.pic = make([][]bool, gpu.w)
	for x := 0; x < gpu.w; x++ {
		gpu.pic[x] = make([]bool, gpu.h)
		for y := 0; y < gpu.h; y++ {
			gpu.pic[x][y] = false
		}
	}
	gpu.color = 0x11FF11 // Some kind of green
}

type CHIP8Sound_i interface {
	init()
	turn_beep() // Change beep status from true to false and from false to true
}

type CHIP8Sound struct {
	turn_on bool
} // Not implemented

func (sound *CHIP8Sound) init() {
	sound.turn_on = false
}

func (sound *CHIP8Sound) turn_beep() {
	sound.turn_on = !sound.turn_on
}

type CHIP8Input_i interface {
	init()
	is_pressed(key uint8) bool
	key_down(key uint8)
	key_up(key uint8)
}

type CHIP8Input struct {
	keys []bool
}

func (input *CHIP8Input) init() {
	input.keys = make([]bool, 16)
}

func (input *CHIP8Input) is_pressed(key uint8) bool {
	return input.keys[key]
}

func (input *CHIP8Input) key_down(key uint8) {
	input.keys[key] = true
}

func (input *CHIP8Input) key_up(key uint8) {
	input.keys[key] = false
}

type CHIP8Console_i interface {
	init()
	loop()
	tick(dt float32)
}

type CHIP8Console struct {
	mem   CHIP8Memory_i
	cpu   CHIP8CPU_i
	gpu   CHIP8GPU_i
	input CHIP8Input_i
	sound CHIP8Sound_i
}

func (console *CHIP8Console) init() {
	console.cpu = new(CHIP8CPU)
	console.mem = new(CHIP8Memory)
	console.gpu = new(CHIP8GPU)
	console.input = new(CHIP8Input)
	console.sound = new(CHIP8Sound)

	console.cpu.init()
	console.mem.init()
	console.gpu.init()
	console.input.init()
	console.sound.init()
}

func (console *CHIP8Console) loop() {

}

func (console *CHIP8Console) tick(dt float32) {

}
