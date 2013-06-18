package main

import (
	"fmt"
)

type Registr uint8
type CPUTimer int8
type OpCode uint16

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
	tick(console *CHIP8Console)
	timer_decrement()
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
	for i := 0; i < 16; i++ {
		cpu.v[i] = 0
	}
	cpu.i = 0
	cpu.pc = 0x200 // First 0x200 byte are interpreter
	cpu.sp = new(Stack)
	cpu.sp.init(16)
	cpu.dt = 0
	cpu.st = 0
}

func (cpu *CHIP8CPU) timer_decrement() {
	cpu.dt--
	cpu.st--
}

func (cpu *CHIP8CPU) tick(console *CHIP8Console) {
	op := OpCode(console.mem.read2(uint32(cpu.pc)))
	cpu.pc += 2
	if cpu.st > 0 {
		cpu.st--
	}
	if cpu.dt > 0 {
		cpu.dt--
	}
	switch uint16(op) & 0xF000 {
	case 0x0000:
		switch uint16(op) & 0xFFFF {
		case 0x00E0:
			cpu.op_00E0(op, console)
		case 0x00EE:
			cpu.op_00EE(op, console)
		default:
			fmt.Printf("Unknown opcode\n")
		}
	case 0x1000:
		cpu.op_1NNN(op, console)
	case 0x2000:
		cpu.op_2NNN(op, console)
	case 0x3000:
		cpu.op_3XNN(op, console)
	case 0x4000:
		cpu.op_4XNN(op, console)
	case 0x5000:
		cpu.op_5XY0(op, console)
	case 0x6000:
		cpu.op_6XNN(op, console)
	case 0x7000:
		cpu.op_7XNN(op, console)
	case 0x8000:
		switch uint16(op) & 0x000F {
		case 0x0:
			cpu.op_8XY0(op, console)
		case 0x1:
			cpu.op_8XY1(op, console)
		case 0x2:
			cpu.op_8XY2(op, console)
		case 0x3:
			cpu.op_8XY3(op, console)
		case 0x4:
			cpu.op_8XY4(op, console)
		case 0x5:
			cpu.op_8XY5(op, console)
		case 0x6:
			cpu.op_8XY6(op, console)
		case 0x7:
			cpu.op_8XY7(op, console)
		case 0xE:
			cpu.op_8XYE(op, console)
		default:
			fmt.Printf("Unknown opcode\n")
		}
	case 0x9000:
		cpu.op_9XY0(op, console)
	case 0xA000:
		cpu.op_ANNN(op, console)
	case 0xB000:
		cpu.op_BNNN(op, console)
	case 0xC000:
		cpu.op_CXNN(op, console)
	case 0xD000:
		cpu.op_DXYN(op, console)
	case 0xE000:
		switch uint16(op) & 0x00FF {
		case 0x009E:
			cpu.op_EX9E(op, console)
		case 0x00A1:
			cpu.op_EXA1(op, console)
		default:
			fmt.Printf("Unknown opcode\n")
		}
	case 0xF000:
		switch uint16(op) & 0x00FF {
		case 0x0007:
			cpu.op_FX07(op, console)
		case 0x000A:
			cpu.op_FX0A(op, console)
		case 0x0015:
			cpu.op_FX15(op, console)

		case 0x0018:
			cpu.op_FX18(op, console)
		case 0x001E:
			cpu.op_FX1E(op, console)
		case 0x0029:
			cpu.op_FX29(op, console)

		case 0x0033:
			cpu.op_FX33(op, console)
		case 0x0055:
			cpu.op_FX55(op, console)
		case 0x0065:
			cpu.op_FX65(op, console)
		default:
			fmt.Printf("Unknown opcode\n")
		}
	default:
		fmt.Printf("Unknown opcode\n")
	}
}
