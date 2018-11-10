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
    fmt.Println("--------Registers-------")
}

func (cpu *CPU) RESET() {
    fmt.Println("[*] RESET Occured")

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

func (cpu *CPU) NMI() {
    // Not Implemented
    // but Partially

    fmt.Println("[*] NMI Occured")

    cpu.StackPush(uint8(cpu.PC >> 0x08)) // PC High
    cpu.StackPush(uint8(cpu.PC & 0x08)) // PC Low
    cpu.StackPush(cpu.P) // Status Register

    cpu.ClearStatusRegister("break")
    cpu.PC = cpu.ReadWord(0xFFFA)
}
func (cpu *CPU) IRQ() {
    // Not Implemented
    // but Partially

    if !cpu.GetStatusRegister("interrupt") {
        fmt.Println("[*] IRQ Occuerd")

        cpu.ClearStatusRegister("break")

        cpu.StackPush(uint8(cpu.PC >> 0x08)) // PC High
        cpu.StackPush(uint8(cpu.PC & 0x08)) // PC Low
        cpu.StackPush(cpu.P) // Status Register

        cpu.SetStatusRegister("interrupt")
        cpu.PC = cpu.ReadWord(0xFFFE)
    }
}
func (cpu *CPU) BRK() {
    // Not Implemented
    // but Partially

    if !cpu.GetStatusRegister("interrupt") {
        fmt.Println("[*] BRK Occuerd")

        cpu.SetStatusRegister("break")
        cpu.SetStatusRegister("interrupt")
        cpu.IncrementPC()

        cpu.StackPush(uint8(cpu.PC >> 0x08)) // PC High
        cpu.StackPush(uint8(cpu.PC & 0x08)) // PC Low
        cpu.StackPush(cpu.P) // Status Register

        cpu.PC = cpu.ReadWord(0xFFFE)
    }
}

func (cpu *CPU) ReadByte(address uint16) uint8 {
    return cpu.Memory[address]
}
func (cpu *CPU) ReadWord(address uint16) uint16 { // Little Endian
    return uint16(cpu.Memory[address]) | (uint16(cpu.Memory[address+1]) << 0x08)
}
func (cpu *CPU) WriteByte(address uint16, val uint8) {
    cpu.Memory[address] = val
}
func (cpu *CPU) WriteWord(address, val uint16) {
    cpu.Memory[address] = uint8(val & 0xff)
    cpu.Memory[address + 1] = uint8(val >> 0x08)
}

func (cpu *CPU) MemoryMapping(val []uint8, start int) {
    for i := 0; i < len(val); i++ {
        cpu.Memory[start + i] = val[i]
    }
}

func (cpu *CPU) Fetch() uint8 {
    defer cpu.IncrementPC() // cannot use statement cpu.Memory[cpu.PC++]
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

func (cpu *CPU) StackPush(val uint8) {
    fmt.Printf("[+] STACK PUSH -> SP : 0x%04X, val : 0x%02X\n", cpu.SP, val)
    cpu.Memory[cpu.SP] = val
    cpu.SP -= 1
}
func (cpu *CPU) StackPop() uint8 {
    //fmt.Println("STACK POP")
    cpu.SP += 1
    return cpu.Memory[cpu.SP]
}

func (cpu *CPU) Run() {

    for ; cpu.PC < 0xFFFF; { // とりあえず一生回しとけ
        opecode := cpu.Fetch()
        fmt.Printf("%02X ", opecode)
    }
    fmt.Println()

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

