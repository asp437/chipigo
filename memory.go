package main

import (
	"io/ioutil"
)

type CHIP8Memory_i interface {
	init()
	read(addr uint32) uint8
	write(addr uint32, val uint8)
	read2(addr uint32) uint16 // Read 2 byte. Special for opcode reading
	read_rom(str string)
}

type CHIP8Memory struct {
	data []uint8
}

func (mem *CHIP8Memory) read(addr uint32) uint8 {
	return mem.data[addr]
}

func (mem *CHIP8Memory) read2(addr uint32) uint16 {
	return (uint16(mem.data[addr]) << 8) | uint16(mem.data[addr+1])
}

func (mem *CHIP8Memory) write(addr uint32, val uint8) {
	mem.data[addr] = val
}

func (mem *CHIP8Memory) read_rom(str string) {
	buffer, err := ioutil.ReadFile(str)
	if err == nil {
		for i := 0; i < len(buffer); i++ {
			mem.data[0x200+i] = uint8(buffer[i])
		}
	}
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
	for i := 0x50; i < 0x1000; i++ {
		mem.data[i] = 0
	}
}
