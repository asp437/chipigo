package main

import "math/rand"

func (cpu *CHIP8CPU) op_0NNN(op OpCode, console *CHIP8Console) {} // 0NNN - Calls RCA 1802 program at address NNN.

func (cpu *CHIP8CPU) op_00E0(op OpCode, console *CHIP8Console) { // 00E0 - Clears the screen.
	console.gpu.clear_screen()
}

func (cpu *CHIP8CPU) op_00EE(op OpCode, console *CHIP8Console) { // 00EE - Returns from a subroutine.
	cpu.pc = uint16(cpu.sp.pop())
}

func (cpu *CHIP8CPU) op_1NNN(op OpCode, console *CHIP8Console) { // 1NNN - Jumps to address NNN.
	n := uint16(op & 0xFFF)
	cpu.pc = n
}

func (cpu *CHIP8CPU) op_2NNN(op OpCode, console *CHIP8Console) { // 2NNN - Calls subroutine at NNN.
	cpu.sp.push(uint32(cpu.pc))
	cpu.pc = uint16(op & 0x0FFF)
}

func (cpu *CHIP8CPU) op_3XNN(op OpCode, console *CHIP8Console) { // 3XNN - Skips the next instruction if VX equals NN.
	x := uint8((op & 0x0F00) >> 16)

	n := Registr(op & 0x00FF)
	if cpu.v[x] == n {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_4XNN(op OpCode, console *CHIP8Console) { // 4XNN - Skips the next instruction if VX doesn't equal NN.
	x := uint8((op & 0x0F00) >> 16)
	n := Registr(op & 0x00FF)
	if cpu.v[x] != n {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_5XY0(op OpCode, console *CHIP8Console) { // 5XY0 - Skips the next instruction if VX equals VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	if cpu.v[x] == cpu.v[y] {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_6XNN(op OpCode, console *CHIP8Console) { // 6XNN - Sets VX to NN.
	x := uint8((op & 0x0F00) >> 16)
	n := Registr(op & 0x00FF)
	cpu.v[x] = n
}

func (cpu *CHIP8CPU) op_7XNN(op OpCode, console *CHIP8Console) { // 7XNN - Adds NN to VX.
	x := uint8((op & 0x0F00) >> 16)
	n := Registr(op & 0x00FF)
	cpu.v[x] += n
}

func (cpu *CHIP8CPU) op_8XY0(op OpCode, console *CHIP8Console) { // 8XY0 - Sets VX to the value of VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	cpu.v[x] = cpu.v[y]
}

func (cpu *CHIP8CPU) op_8XY1(op OpCode, console *CHIP8Console) { // 8XY1 - Sets VX to VX or VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	cpu.v[x] = cpu.v[x] | cpu.v[y]
}

func (cpu *CHIP8CPU) op_8XY2(op OpCode, console *CHIP8Console) { // 8XY2 - Sets VX to VX and VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	cpu.v[x] = cpu.v[x] & cpu.v[y]
}

func (cpu *CHIP8CPU) op_8XY3(op OpCode, console *CHIP8Console) { // 8XY3 - Sets VX to VX xor VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	cpu.v[x] = cpu.v[x] ^ cpu.v[y]
}

func (cpu *CHIP8CPU) op_8XY4(op OpCode, console *CHIP8Console) { // 8XY4 - Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	sum := uint32(cpu.v[x] + cpu.v[y])
	if sum > 0xFFFF {
		cpu.v[0xF] = 1
	} else {
		cpu.v[0xF] = 0
	}
	cpu.v[x] = Registr(sum & 0xFFFF)
}

func (cpu *CHIP8CPU) op_8XY5(op OpCode, console *CHIP8Console) { // 8XY5 - VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	if cpu.v[x] > cpu.v[y] {
		cpu.v[0xF] = 1
	} else {
		cpu.v[0xF] = 0
	}
	cpu.v[x] = Registr(cpu.v[x] - cpu.v[y])
}

func (cpu *CHIP8CPU) op_8XY6(op OpCode, console *CHIP8Console) { // 8XY6 - Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.[2]
	x := uint8((op & 0x0F00) >> 16)
	cpu.v[0xF] = cpu.v[x] & 0x0001
	cpu.v[x] = cpu.v[x] >> 1
}

func (cpu *CHIP8CPU) op_8XY7(op OpCode, console *CHIP8Console) { // 8XY7 - Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	if cpu.v[y] > cpu.v[x] {
		cpu.v[0xF] = 1
	} else {
		cpu.v[0xF] = 0
	}
	cpu.v[x] = Registr(cpu.v[y] - cpu.v[x])
}

func (cpu *CHIP8CPU) op_8XYE(op OpCode, console *CHIP8Console) { // 8XYE -  Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.[2]
	x := uint8((op & 0x0F00) >> 16)
	cpu.v[0xF] = (cpu.v[x] & 0x80) >> 15
	cpu.v[x] = cpu.v[x] << 1
}

func (cpu *CHIP8CPU) op_9XY0(op OpCode, console *CHIP8Console) { // 9XY0 -  Skips the next instruction if VX doesn't equal VY.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	if cpu.v[x] != cpu.v[y] {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_ANNN(op OpCode, console *CHIP8Console) { // ANNN -  Sets I to the address NNN.
	cpu.i = uint16(op & 0x0FFF)
}

func (cpu *CHIP8CPU) op_BNNN(op OpCode, console *CHIP8Console) { // BNNN -  Jumps to the address NNN plus V0.
	cpu.pc = uint16(op&0x0FFF) + uint16(cpu.v[0x0])
}

func (cpu *CHIP8CPU) op_CXNN(op OpCode, console *CHIP8Console) { // CXNN -  Sets VX to a random number and NN.
	x := uint8((op & 0x0F00) >> 16)
	n := uint8(op & 0x00FF)
	cpu.v[x] = Registr(uint8(rand.Int()) & n)
}

func (cpu *CHIP8CPU) op_DXYN(op OpCode, console *CHIP8Console) { // DXYN -  Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels. Each row of 8 pixels is read as bit-coded (with the most significant bit of each byte displayed on the left) starting from memory location I; I value doesn't change after the execution of this instruction. As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that doesn't happen.
	x := uint8((op & 0x0F00) >> 16)
	y := uint8((op & 0x00F0) >> 8)
	n := uint8(op & 0x00FF)
	for i := 0; i < int(n); i++ {
		console.gpu.draw_line8(uint8(cpu.v[x]), uint8(cpu.v[y])+uint8(i), console.mem.read(uint32(cpu.i)+uint32(i)))
	}
}

func (cpu *CHIP8CPU) op_EX9E(op OpCode, console *CHIP8Console) { // EX9E -  Skips the next instruction if the key stored in VX is pressed.
	x := uint8((op & 0x0F00) >> 16)
	if console.input.is_pressed(uint8(cpu.v[x])) {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_EXA1(op OpCode, console *CHIP8Console) { // EXA1 -  Skips the next instruction if the key stored in VX isn't pressed.
	x := uint8((op & 0x0F00) >> 16)
	if !console.input.is_pressed(uint8(cpu.v[x])) {
		cpu.pc += 2
	}
}

func (cpu *CHIP8CPU) op_FX07(op OpCode, console *CHIP8Console) { // FX07 -  Sets VX to the value of the delay timer.
	x := uint((op & 0x0F00) >> 16)
	cpu.v[x] = Registr(cpu.dt)
}

func (cpu *CHIP8CPU) op_FX0A(op OpCode, console *CHIP8Console) { // FX0A -  A key press is awaited, and then stored in VX.
	x := uint((op & 0x0F00) >> 16)
	for i := 0; i < 0x10; i++ {
		if console.input.is_pressed(uint8(i)) {
			cpu.v[x] = Registr(i)
			return
		}
	}
	cpu.pc -= 2 // If no key pressed, try on next tick
}

func (cpu *CHIP8CPU) op_FX15(op OpCode, console *CHIP8Console) { // FX15 -  Sets the delay timer to VX.
	x := uint8((op & 0x0F00) >> 16)
	cpu.dt = CPUTimer(cpu.v[x])
}

func (cpu *CHIP8CPU) op_FX18(op OpCode, console *CHIP8Console) { // FX18 -  Sets the sound timer to VX.
	x := uint8((op & 0x0F00) >> 16)
	cpu.st = CPUTimer(cpu.v[x])
}

func (cpu *CHIP8CPU) op_FX1E(op OpCode, console *CHIP8Console) { // FX1E -  Adds VX to I.[3]
	// Note: VF is set to 1 when range overflow (I+VX>0xFFF), and 0 when there isn't.
	// This is undocumented feature of the Chip-8 and used by Spacefight 2019! game.
	x := uint16((op & 0x0F00) >> 16)
	cpu.i += x
	if cpu.i > 0xFFF {
		cpu.v[0xF] = 1
	} else {
		cpu.v[0xF] = 0
	}
}

func (cpu *CHIP8CPU) op_FX29(op OpCode, console *CHIP8Console) { // FX29 -  Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
	x := uint8((op & 0x0F00) >> 16)
	cpu.i = uint16(cpu.v[x] * 0x5)
}

func (cpu *CHIP8CPU) op_FX33(op OpCode, console *CHIP8Console) { // FX33 -  Stores the Binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2. (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
	x := uint8((op & 0x0F00) >> 16)
	var a, b, c uint16
	a = uint16(cpu.v[x]) % 10
	b = (uint16(cpu.v[x]%100) - a) / 10
	c = ((uint16(cpu.v[x]) % 1000) - a - b*10) / 100
	console.mem.write(uint32(cpu.i), uint8(c))
	console.mem.write(uint32(cpu.i)+1, uint8(b))
	console.mem.write(uint32(cpu.i)+2, uint8(a))
}

func (cpu *CHIP8CPU) op_FX55(op OpCode, console *CHIP8Console) { // FX55 -  Stores V0 to VX in memory starting at address I.[4]
	x := uint8((op & 0x0F00) >> 16)
	var i uint8
	for i = 0; i <= x; i++ {
		console.mem.write(uint32(cpu.i)+uint32(i), uint8(cpu.v[uint8(i)]))
	}
}

func (cpu *CHIP8CPU) op_FX65(op OpCode, console *CHIP8Console) { // FX65 -  Fills V0 to VX with values from memory starting at address I.[4]
	x := uint8((op & 0x0F00) >> 16)
	var i uint8
	for i = 0; i <= x; i++ {
		cpu.v[uint8(i)] = Registr(console.mem.read(uint32(cpu.i) + uint32(i)))
	}
}

/*
Legend:
[ ] - not implemented
[\] - implemented, no tests
[V] - implemented with tests
[B] - implemented with bugs

       Opcode	Explanation
[X]    0NNN		Calls RCA 1802 program at address NNN. (No implement needed)
[\]    00E0		Clears the screen.
[\]    00EE		Returns from a subroutine.
[\]    1NNN		Jumps to address NNN.
[\]    2NNN		Calls subroutine at NNN.
[\]    3XNN		Skips the next instruction if VX equals NN.
[\]    4XNN		Skips the next instruction if VX doesn't equal NN.
[\]    5XY0		Skips the next instruction if VX equals VY.
[\]    6XNN		Sets VX to NN.
[\]    7XNN		Adds NN to VX.
[\]    8XY0		Sets VX to the value of VY.
[\]    8XY1		Sets VX to VX or VY.
[\]    8XY2		Sets VX to VX and VY.
[\]    8XY3		Sets VX to VX xor VY.
[\]    8XY4		Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
[\]    8XY5		VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
[\]    8XY6		Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.[2]
[\]    8XY7		Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
[\]    8XYE		Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.[2]
[\]    9XY0		Skips the next instruction if VX doesn't equal VY.
[\]    ANNN		Sets I to the address NNN.
[\]    BNNN		Jumps to the address NNN plus V0.
[\]    CXNN		Sets VX to a random number and NN.
[\]    DXYN		Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels. Each row of 8 pixels is read as bit-coded (with the most significant bit of each byte displayed on the left) starting from memory location I; I value doesn't change after the execution of this instruction. As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that doesn't happen.
[\]    EX9E		Skips the next instruction if the key stored in VX is pressed.
[\]    EXA1		Skips the next instruction if the key stored in VX isn't pressed.
[\]    FX07		Sets VX to the value of the delay timer.
[\]    FX0A		A key press is awaited, and then stored in VX.
[\]    FX15		Sets the delay timer to VX.
[\]    FX18		Sets the sound timer to VX.
[\]    FX1E		Adds VX to I.[3]
[\]    FX29		Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
[\]    FX33		Stores the Binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2. (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
[\]    FX55		Stores V0 to VX in memory starting at address I.[4]
[\]    FX65		Fills V0 to VX with values from memory starting at address I.[4]
*/
