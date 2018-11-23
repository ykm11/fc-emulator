package processor

import (
    "fmt"
)

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

func (cpu *CPU) MemoryMapping(val []uint8, start int) {
    if(start + len(val) > 0x10000){
        panic("MemoryMapping Error")
    }
    for i := 0; i < len(val); i++ {
        cpu.Memory[start + i] = val[i]
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
