package testdata

import (
    "testing"
    "../processor"
)


func TestStatusRegister(t *testing.T) {
    cpu := processor.CPU{}

    if cpu.P != 0x00 {
        t.Fatal("")
    }

    if cpu.GetStatusRegister("reserved") {
        t.Fatal("Reserved bit being reset")
    }
    if cpu.GetStatusRegister("C") {
        t.Fatal("C Should be clear")
    }
    if cpu.GetStatusRegister("Z") {
        t.Fatal("Z Should be clear")
    }
    if cpu.GetStatusRegister("D") {
        t.Fatal("D Should be clear")
    }
    if cpu.GetStatusRegister("B") {
        t.Fatal("B Should be clear")
    }
    if cpu.GetStatusRegister("V") {
        t.Fatal("V Should be clear")
    }
    if cpu.GetStatusRegister("N") {
        t.Fatal("N Should be clear")
    }
}

func TestRegisterSetAndClear(t *testing.T) {
    cpu := processor.CPU{}
    if cpu.P != 0x00 {
        t.Fatal("P register should be 0x00")
    }

    cpu.SetStatusRegister("N")
    if cpu.P != 0x80 {
        t.Fatal("N is set and others should be clear")
    }
    cpu.ClearStatusRegister("N")
    if cpu.P == 0x80 {
        t.Fatal("N should be clear")
    }

    cpu.SetStatusRegister("C")
    if cpu.P != 0x01 {
        t.Fatal("C is set and others should be clear")
    }
    cpu.ClearStatusRegister("C")
    if cpu.P == 0x01 {
        t.Fatal("C should be clear")
    }

    cpu.SetStatusRegister("D")
    cpu.SetStatusRegister("B")
    if cpu.P != 0x18 {
        t.Fatal("P should be 0x18")
    }
    cpu.ClearStatusRegister("D")
    if cpu.P != 0x10 {
        t.Fatal("P should be 0x10")
    }
}

func TestCpuReset(t *testing.T) {
    cpu := processor.CPU{}
    cpu.RESET()
    if !cpu.GetStatusRegister("reserved") {
        t.Fatal("Reserved flag should be always set")
    }
}

func TestMemoryWrite(t *testing.T) {
    cpu := processor.CPU{}

    cpu.WriteByte(0x1000, 0x11)
    if cpu.ReadByte(0x1000) != 0x11 {
        t.Fatal("Memory[0x1000] should be 0x11")
    }

    cpu.WriteWord(0x1144, 0x4567)
    if cpu.ReadWord(0x1144) != 0x4567 {
        t.Fatal("Memory[0x1144, 0x1145] should be 0x67 0x45 respectively")
    }
    if cpu.ReadByte(0x1144) != 0x67 {
        t.Fatal("Memory[0x1144] should be 0x67")
    }
    if cpu.ReadByte(0x1145) != 0x45 {
        t.Fatal("Memory[0x1145] should be 0x45")
    }
}
