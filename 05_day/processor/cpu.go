package processor

import (
    "fmt"
)

type CPU struct {
    Registers
    Memory [0x10000]uint8
}


func (cpu *CPU) RESET() {
    fmt.Println("[*] RESET Occured")

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

func (cpu *CPU) Run() {
    opecodes := GetInstructions()

    for ; cpu.PC < 0xFFFA; { // とりあえず一生回しとけ
        pc := cpu.PC
        code := cpu.Fetch()
        fmt.Printf("[+] PC, opecode : 0x%X, %X\n", pc, code)

        opecode := opecodes[code]
        switch opecode.mode { // Partially
        case "accumulator":
            fmt.Println("[*] Accumulator")

        case "immediate":
            fmt.Println("[*] Immidiate")
            data := cpu.Fetch()
            fmt.Printf("[+] data : %X\n", data)

        case "absolute":
            fmt.Println("[*] Absolute")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            valid_add := uint16(high_address) << 0x08 | uint16(low_address)
            fmt.Printf("[+] High, Low : 0x%X 0x%X\n", high_address, low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", valid_add)

        case "absoluteX":
            fmt.Println("[*] AbsoluteX")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            valid_add := uint16(high_address) << 0x08 | uint16(low_address) + uint16(cpu.X)
            fmt.Printf("[+] Valid Address : 0x%X\n", valid_add)

        case "absoluteY":
            fmt.Println("[*] AbsoluteY")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            valid_add := uint16(high_address) << 0x08 | uint16(low_address) + uint16(cpu.Y)
            fmt.Printf("[+] Valid Address : 0x%X\n", valid_add)

        case "zeroPage":
            fmt.Println("[*] ZeroPage")
            low_address := cpu.Fetch()
            fmt.Printf("[+] Low Address : 0x%X\n", low_address)

        case "zeroPageX":
            fmt.Println("[*] ZeroPageX")
            low_address := cpu.Fetch() + cpu.X
            fmt.Printf("[+] Low Address : 0x%X\n", low_address)

        case "zeroPageY":
            fmt.Println("[*] ZeroPageY")
            low_address := cpu.Fetch() + cpu.Y
            fmt.Printf("[+] Low Address : 0x%X\n", low_address)

        case "implied":
            fmt.Println("[*] Implied")
            // No address specification

        case "relative":
            fmt.Println("[*] Relative")
            offset := int8(cpu.Fetch())
            fmt.Printf("[+] offset : %d\n", offset)

        case "indexedIndirectX":
            fmt.Println("[*] Indirect X")
            low_address := cpu.Fetch()
            fmt.Println(low_address)

        case "indirectIndexedY":
            fmt.Println("[*] Indirect Y")
            low_address := cpu.Fetch()
            fmt.Println(low_address)

        case "absoluteIndirect":
            fmt.Println("[*] Absolute Indirect")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            valid_add := uint16(high_address << 0x08) | uint16(low_address)
            fmt.Println(valid_add)

        default :
            fmt.Println("[*] Nothing to do")
        }
        fmt.Println()
    }
}

