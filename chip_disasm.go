package main

import (
	"fmt"
)

func disasm_rom(rom_path string) {
	mem := new(CHIP8Memory)
	mem.init()
	mem.read_rom(rom_path)
	pc := 0x200
	for {
		op := OpCode(mem.read2(uint32(pc)))
		pc += 2
		if op == 0x0000 {
			break
		}
		switch uint16(op) & 0xF000 {
		case 0x0000:
			switch uint16(op) & 0xFFFF {
			case 0x00E0:
				fmt.Printf("CLS\n")
			case 0x00EE:
				fmt.Printf("RET\n")
			default:
				fmt.Printf("SYS %X\n", (op & 0x0FFF))
			}
		case 0x1000:
			fmt.Printf("JP %X\n", (op & 0x0FFF))
		case 0x2000:
			fmt.Printf("CALL %X\n", (op & 0x0FFF))
		case 0x3000:
			fmt.Printf("SE V%x, %X\n", (op&0x0F00)>>8, (op & 0x00FF))
		case 0x4000:
			fmt.Printf("SNE V%x, %X\n", (op&0x0F00)>>8, (op & 0x00FF))
		case 0x5000:
			fmt.Printf("SE V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
		case 0x6000:
			fmt.Printf("LD V%x, %X\n", (op&0x0F00)>>8, (op & 0x00FF))
		case 0x7000:
			fmt.Printf("ADD V%x, %X\n", (op&0x0F00)>>8, (op & 0x00FF))
		case 0x8000:
			switch uint16(op) & 0x000F {
			case 0x0:
				fmt.Printf("LD V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x1:
				fmt.Printf("OR V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x2:
				fmt.Printf("AND V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x3:
				fmt.Printf("XOR V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x4:
				fmt.Printf("ADD V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x5:
				fmt.Printf("SUB V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x6:
				fmt.Printf("SHR V%x {, V%x}\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0x7:
				fmt.Printf("SUBN V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			case 0xE:
				fmt.Printf("SHL V%x {, V%x}\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
			default:
				fmt.Printf("Unknown opcode\n")
			}
		case 0x9000:
			fmt.Printf("SNE V%x, V%x\n", (op&0x0F00)>>8, (op&0x00F0)>>4)
		case 0xA000:
			fmt.Printf("LD I, %X\n", (op & 0x0FFF))
		case 0xB000:
			fmt.Printf("JP V0, %X\n", (op & 0x0FFF))
		case 0xC000:
			fmt.Printf("RND V%x, %X\n", (op&0x0F00)>>8, (op & 0x00FF))
		case 0xD000:
			fmt.Printf("DRW V%x, V%x, %X\n", (op&0x0F00)>>8, (op&0x00F0)>>4, (op & 0x000F))
		case 0xE000:
			switch uint16(op) & 0x00FF {
			case 0x009E:
				fmt.Printf("SKP V%x\n", (op&0x0F00)>>8)
			case 0x00A1:
				fmt.Printf("SKNP V%x\n", (op&0x0F00)>>8)
			default:
				fmt.Printf("Unknown opcode\n")
			}
		case 0xF000:
			switch uint16(op) & 0x00FF {
			case 0x0007:
				fmt.Printf("LD V%x, DT\n", (op&0x0F00)>>8)
			case 0x000A:
				fmt.Printf("LD V%x, K\n", (op&0x0F00)>>8)
			case 0x0015:
				fmt.Printf("LD DT, V%x\n", (op&0x0F00)>>8)
			case 0x0018:
				fmt.Printf("LD ST, V%x\n", (op&0x0F00)>>8)
			case 0x001E:
				fmt.Printf("ADD I, V%x\n", (op&0x0F00)>>8)
			case 0x0029:
				fmt.Printf("LD F, V%x\n", (op&0x0F00)>>8)
			case 0x0033:
				fmt.Printf("LD B, V%x\n", (op&0x0F00)>>8)
			case 0x0055:
				fmt.Printf("LD [I], V%x\n", (op&0x0F00)>>8)
			case 0x0065:
				fmt.Printf("LD V%x, [I]\n", (op&0x0F00)>>8)
			default:
				fmt.Printf("Unknown opcode\n")
			}
		default:
			fmt.Printf("Unknown opcode\n")
		}
	}
}
