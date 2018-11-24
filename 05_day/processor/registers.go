package processor

import (
    "fmt"
)

type Registers struct {
    A uint8     // Accumulator
    X uint8     // Index Register
    Y uint8     // Index Register
    P uint8     // Status Register
    SP uint16   // Stack Register
    PC uint16   // Program Counter
}

func (cpu *CPU) ShowRegister() { // For debugging
    fmt.Println("\n--------Registers-------")
    fmt.Printf("[+] Register A  : 0x%02x\n", cpu.A)
    fmt.Printf("[+] Rsgister X  : 0x%02x\n", cpu.X)
    fmt.Printf("[+] Rsgister Y  : 0x%02x\n", cpu.Y)
    fmt.Printf("[+] Rsgister P  : 0b%08b\n", cpu.P)
    fmt.Printf("[+] Rsgister SP : 0x%04x\n", cpu.SP)
    fmt.Printf("[+] Rsgister PC : 0x%04x\n", cpu.PC)
    fmt.Println("--------Registers-------")
}

func (cpu *CPU) GetStatusRegister(q string) bool {
    switch q {
    case "carry", "C":
        return (cpu.Registers.P & 0x01) != 0
    case "zero", "Z":
        return (cpu.Registers.P & 0x02) != 0
    case "interrupt", "I":
        return (cpu.Registers.P & 0x04) != 0
    case "decimal", "D":
        return (cpu.Registers.P & 0x08) != 0
    case "break", "B":
        return (cpu.Registers.P & 0x10) != 0
    case "reserved":
        return (cpu.Registers.P & 0x20) != 0
    case "overflow", "V":
        return (cpu.Registers.P & 0x40) != 0
    case "negative", "N":
        return (cpu.Registers.P & 0x80) != 0
    default:
        panic("Unknown query")
    }
}

func (cpu *CPU) SetStatusRegister(q string) {
    switch q {
    case "carry", "C":
        cpu.P |= 0x01
    case "zero", "Z":
        cpu.P |= 0x02
    case "interrupt", "I":
        cpu.P |= 0x04
    case "decimal", "D":
        cpu.P |= 0x08
    case "break", "B":
        cpu.P |= 0x10
    case "reserved":
        panic("Forbidden")
    case "overflow", "V":
        cpu.P |= 0x40
    case "negative", "N":
        cpu.P |= 0x80
    default:
        panic("Unknown query")
    }
}

func (cpu *CPU) ClearStatusRegister(q string) {
    switch q {
    case "carry", "C":
        cpu.P &= 0xFE
    case "zero", "Z":
        cpu.P &= 0xFD
    case "interrupt", "I":
        cpu.P &= 0xFC
    case "decimal", "D":
        cpu.P &= 0xF7
    case "break", "B":
        cpu.P &= 0xEF
    case "reserved":
        panic("Forbidden")
    case "overflow", "V":
        cpu.P &= 0xCF
    case "negative", "N":
        cpu.P &= 0x7F
    default:
        panic("Unknown query")
    }
}

func (cpu *CPU) PopPC() uint16 {
    high_address := cpu.StackPop()
    low_address := cpu.StackPop()
    fmt.Printf("[+] POP PC (high, low) : %X, %X\n", high_address, low_address)
    return uint16(high_address) << 0x08 | uint16(low_address)
}

func (cpu *CPU) PushPC(){
    fmt.Printf("[+] PUSH PC : %X\n", cpu.PC)
    cpu.StackPush(uint8(cpu.PC & 0xFF))
    cpu.StackPush(uint8(cpu.PC >> 0x08))
}
