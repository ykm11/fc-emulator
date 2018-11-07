package processor

import (
    "fmt"
)

type Registers struct {
    A uint8   // Accumulator
    X uint8   // Index Register
    Y uint8   // Index Register
    P uint8   // Status Register
    SP uint16   // Stack Register
    PC uint16 // Program Counter
}

type CPU struct {
    Registers
    Memory [0x10000]uint8
}


func (cpu *CPU) ShowRegister() { // For debugging
    fmt.Println("\n--------Registers-------")
    fmt.Printf("[+] Register A  : 0x%02x\n", cpu.A)
    fmt.Printf("[+] Rsgister X  : 0x%02x\n", cpu.X)
    fmt.Printf("[+] Rsgister Y  : 0x%02x\n", cpu.Y)
    fmt.Printf("[+] Rsgister P  : 0b%08b\n", cpu.P)
    fmt.Printf("[+] Rsgister SP : 0x%04x\n", cpu.SP)
    fmt.Printf("[+] Rsgister PC : 0x%04x\n", cpu.PC)
    fmt.Println("--------Registers-------\n")
}

func (cpu *CPU) Reset() {
    fmt.Println("[*] Register Reset")

    cpu.Memory[0xFFFC] = 0x00
    cpu.Memory[0xFFFD] = 0x80

    // Initialize Registers (default)
    cpu.A = 0x00
    cpu.X = 0x00
    cpu.Y = 0x00
    cpu.P = 0x39
    cpu.SP = 0x01FD
    cpu.PC = cpu.ReadWord(0xFFFC)
}

func (cpu *CPU) Nmi() {
}
func (cpu *CPU) Irq() {
}
func (cpu *CPU) Brk() {
}

func (cpu *CPU) ReadByte(address uint16) uint8 {
    return cpu.Memory[address]
}
func (cpu *CPU) ReadWord(address uint16) uint16 { // Little Endian
    return uint16(cpu.Memory[address]) | (uint16(cpu.Memory[address+1]) << 0x08)
}

func (cpu *CPU) MemoryMapping(val []uint8, start int) {
    for i := 0; i < len(val); i++ {
        cpu.Memory[start + i] = val[i]
    }
}

func (cpu *CPU) Fetch() uint8 {
    defer cpu.IncrementPC() // cannot use statement [cpu.PC++]
    return cpu.Memory[cpu.PC]
}

func (cpu *CPU) IncrementPC() {
    if cpu.PC < 0xFFFF {
        cpu.PC += 1
    } else {
        // Raise panic for now
        panic("End of Memory")
    }
}


func (cpu CPU) GetStatusRegister(q string) bool {
    switch q {
    case "carry":
        return (cpu.Registers.P & 0x01) != 0
    case "zero":
        return (cpu.Registers.P & 0x02) != 0
    case "interrupt":
        return (cpu.Registers.P & 0x04) != 0
    case "decimal":
        return (cpu.Registers.P & 0x08) != 0
    case "break":
        return (cpu.Registers.P & 0x10) != 0
    case "reserved":
        return (cpu.Registers.P & 0x20) != 0
    case "overflow":
        return (cpu.Registers.P & 0x40) != 0
    case "negative":
        return (cpu.Registers.P & 0x80) != 0
    default:
        panic("Unknown query")
    }
}
