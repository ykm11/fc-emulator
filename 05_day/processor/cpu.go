package processor

import (
    "fmt"
    "../utils"
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
    var addrOrData uint16

    for ; cpu.PC < 0xFFFE ; { // とりあえず一生回しとけ
        pc := cpu.PC
        code := cpu.Fetch()
        fmt.Printf("[+] PC, opecode : 0x%X, %X\n", pc, code)

        opecode := opecodes[code]
        switch opecode.mode { // Partially
        case "accumulator":
            //fmt.Println("[*] Accumulator")
            addrOrData = 0

        case "immediate":
            //fmt.Println("[*] Immidiate")
            addrOrData = uint16(cpu.Fetch())
            fmt.Printf("[+] data : %X\n", addrOrData)

        case "absolute":
            //fmt.Println("[*] Absolute")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            addrOrData = uint16(high_address) << 0x08 | uint16(low_address)
            fmt.Printf("[+] High, Low : 0x%X 0x%X\n", high_address, low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "absoluteX":
            //fmt.Println("[*] AbsoluteX")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            addrOrData := uint16(high_address) << 0x08 | uint16(low_address) + uint16(cpu.X)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "absoluteY":
            //fmt.Println("[*] AbsoluteY")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            addrOrData := uint16(high_address) << 0x08 | uint16(low_address) + uint16(cpu.Y)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "zeroPage":
            //fmt.Println("[*] ZeroPage")
            low_address := cpu.Fetch()
            addrOrData = uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "zeroPageX":
            //fmt.Println("[*] ZeroPageX")
            low_address := cpu.Fetch() + cpu.X
            addrOrData = uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "zeroPageY":
            //fmt.Println("[*] ZeroPageY")
            low_address := cpu.Fetch() + cpu.Y
            addrOrData = uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "implied":
            //fmt.Println("[*] Implied")
            // No address specification
            addrOrData = 0

        case "relative":
            //fmt.Println("[*] Relative")
            offset := cpu.Fetch()
            addrOrData = uint16(offset)
            fmt.Printf("[+] offset : %X\n", addrOrData)

        case "indexedIndirectX":
            //fmt.Println("[*] Indirect X")
            low_address := cpu.Fetch() + cpu.X
            addrOrData = uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "indirectIndexedY":
            //fmt.Println("[*] Indirect Y")
            low_address := cpu.Fetch() + cpu.Y
            addrOrData = uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        case "absoluteIndirect":
            //fmt.Println("[*] Absolute Indirect")
            low_address := cpu.Fetch()
            high_address := cpu.Fetch()
            addrOrData = uint16(high_address << 0x08) | uint16(low_address)
            fmt.Printf("[+] Valid Address : 0x%X\n", addrOrData)

        default :
            fmt.Println("[*] Not Implemented Instruction\n")
            continue
        }
        cpu.ExecInstruction(opecode.syntax, opecode.mode, addrOrData)
    }
}

func (cpu *CPU) ExecInstruction(syntax, mode string, addrOrData uint16) {
    fmt.Printf("[+] Sybtax, Mode, AddrOrData : %s, %s, %X\n\n", syntax, mode, addrOrData)

    switch syntax {
    // ロード
    case "LDA":
        if mode == "immediate" {
            cpu.A = uint8(addrOrData & 0xFF)
        } else {
            cpu.A = cpu.ReadByte(addrOrData)
        }

        if cpu.A & 0x80 == 0 {
            cpu.SetStatusRegister("N")
        } else {
            cpu.ClearStatusRegister("N")
        }

        if cpu.A == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }

    case "LDX":
        if mode == "immediate" {
            cpu.X = uint8(addrOrData & 0xFF)
        } else {
            cpu.X = cpu.ReadByte(addrOrData)
        }

        if cpu.X & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }

        if cpu.X == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }

    case "LDY":
        if mode == "immediate" {
            cpu.Y = uint8(addrOrData & 0xFF)
        } else {
            cpu.Y = cpu.ReadByte(addrOrData)
        }

        if cpu.Y & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }

        if cpu.Y == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }


    // ストア
    case "STA":
        cpu.WriteByte(addrOrData, cpu.A)

    case "STX":
        cpu.WriteByte(addrOrData, cpu.X)

    case "STY":
        cpu.WriteByte(addrOrData, cpu.Y)

    // 比較
    case "CMP":
        var data uint8
        if mode == "immediate" {
            data = uint8(addrOrData & 0xFF)
        } else {
            data = cpu.ReadByte(addrOrData)
        }

        cmp := utils.Compare(cpu.A, data)
        fmt.Printf("[+] A, M, cmp : %d, %d, %d\n", cpu.A, data, cmp)
        if cmp == -1 {
            cpu.SetStatusRegister("N")
            cpu.ClearStatusRegister("C")
            cpu.ClearStatusRegister("Z")
        } else if cmp == 1 {
            cpu.SetStatusRegister("C")
            cpu.ClearStatusRegister("N")
            cpu.ClearStatusRegister("Z")
        } else {
            cpu.SetStatusRegister("C")
            cpu.SetStatusRegister("Z")
            cpu.ClearStatusRegister("N")
        }

    case "CPX":
        var data uint8
        if mode == "immediate" {
            data = uint8(addrOrData & 0xFF)
        } else {
            data = cpu.ReadByte(addrOrData)
        }

        cmp := utils.Compare(cpu.X, data)
        fmt.Printf("[+] X, M, cmp : %d, %d, %d\n", cpu.X, data, cmp)
        if cmp == -1 {
            cpu.SetStatusRegister("N")
            cpu.ClearStatusRegister("C")
            cpu.ClearStatusRegister("Z")
        } else if cmp == 1 {
            cpu.SetStatusRegister("C")
            cpu.ClearStatusRegister("N")
            cpu.ClearStatusRegister("Z")
        } else {
            cpu.SetStatusRegister("C")
            cpu.SetStatusRegister("Z")
            cpu.ClearStatusRegister("N")
        }

    case "CPY":
        var data uint8
        if mode == "immediate" {
            data = uint8(addrOrData & 0xFF)
        } else {
            data = cpu.ReadByte(addrOrData)
        }

        cmp := utils.Compare(cpu.Y, data)
        fmt.Printf("[+] Y, M, cmp : %d, %d, %d\n", cpu.Y, data, cmp)
        if cmp == -1 {
            cpu.SetStatusRegister("N")
            cpu.ClearStatusRegister("C")
            cpu.ClearStatusRegister("Z")
        } else if cmp == 1 {
            cpu.SetStatusRegister("C")
            cpu.ClearStatusRegister("N")
            cpu.ClearStatusRegister("Z")
        } else {
            cpu.SetStatusRegister("C")
            cpu.SetStatusRegister("Z")
            cpu.ClearStatusRegister("N")
        }

    // 条件分岐
    case "BCC":
        if !cpu.GetStatusRegister("C") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }
    case "BCS":
        if cpu.GetStatusRegister("C") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }

    case "BNE":
        if !cpu.GetStatusRegister("Z") {
            fmt.Printf("[+] BNE PC : %X\n", cpu.PC)
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
            fmt.Printf("[+] BNE PC : %X\n", cpu.PC)
        }
    case "BEQ":
        if cpu.GetStatusRegister("Z") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }

    case "BVC":
        if !cpu.GetStatusRegister("V") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }
    case "BVS":
        if cpu.GetStatusRegister("V") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }

    case "BPL":
        if !cpu.GetStatusRegister("N") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }
    case "BMI":
        if cpu.GetStatusRegister("N") {
            if addrOrData & 0x80 == 0 {
                cpu.PC += (addrOrData & 0xFF)
            } else {
                cpu.PC -= (0x100 - (addrOrData & 0xFF))
            }
        }

    // インクリメント, デクリメント
    case "INC":
        fmt.Printf("\t[+] INC Address, data : %X, %X\n", addrOrData, cpu.ReadByte(addrOrData))
        data := (cpu.ReadByte(addrOrData) + 1) & 0xFF
        if data & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if data == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }
        cpu.WriteByte(addrOrData, data)
    case "DEC":
        fmt.Printf("\t[+] DEC Address, data : %X, %X\n", addrOrData, cpu.ReadByte(addrOrData))
        data := (cpu.ReadByte(addrOrData) - 1) & 0xFF
        if data & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if data == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }
        cpu.WriteByte(addrOrData, data)

    case "INX":
        cpu.X += 1
        if cpu.X & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if cpu.X == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }
    case "DEX":
        cpu.X -= 1
        if cpu.X & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if cpu.X == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }

    case "INY":
        cpu.Y += 1
        if cpu.X & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if cpu.Y == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }
    case "DEY":
        cpu.Y -= 1
        if cpu.Y & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }
        if cpu.Y == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }

    // フラグ操作
    case "CLC":
        cpu.ClearStatusRegister("C")
    case "SEC":
        cpu.SetStatusRegister("C")

    case "CLI":
        cpu.ClearStatusRegister("I")
    case "SEI":
        cpu.SetStatusRegister("I")

    case "CLD":
        cpu.ClearStatusRegister("D")
    case "SED":
        cpu.SetStatusRegister("D")

    case "CLV":
        cpu.ClearStatusRegister("V")

    // スタック操作
    case "PHA":
        cpu.StackPush(cpu.A)
    case "PLA":
        cpu.A = cpu.StackPop()

        if cpu.A & 0x80 == 0 {
            cpu.ClearStatusRegister("N")
        } else {
            cpu.SetStatusRegister("N")
        }

        if cpu.A == 0 {
            cpu.SetStatusRegister("Z")
        } else {
            cpu.ClearStatusRegister("Z")
        }

    case "PHP":
        cpu.StackPush(cpu.P)
        cpu.SetStatusRegister("B")
    case "PLP":
        cpu.P = cpu.StackPop()

    // ジャンプ命令
    case "JMP":
        cpu.PC = addrOrData

    case "JSR":
        cpu.PushPC()
        cpu.PC = addrOrData

    case "RTS":
        cpu.PC = cpu.PopPC()
        fmt.Printf("[+] Returned PC : %X\n", cpu.PC)

    // 割り込み
    case "BRK":
        cpu.BRK()

    // 何もない
    case "NOP":
    default :
        return
    }
}
