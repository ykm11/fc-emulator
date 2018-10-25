package main

import (
    "fmt"
    "os"
    "unsafe"
    "bytes"
    "encoding/binary"
)

type Header struct { // 16 bytes
    Constant [4]byte
    PRG_ROM_SIZE uint8
    CHR_ROM_SIZE uint8
    Flags6 uint8
    Flags7 uint8
    PRG_RAM_SIZE uint8
    Flags9 uint8
    Flags10 uint8
    _ [5]byte
}

type Register struct {
    A uint8 // Accumulator
    X uint8 // Index Register
    Y uint8 // Index Register
    S uint8 // Stack Pointer
    P uint8 // Status Register
    PC uint16 // Program Counter
}


func readHeader(file *os.File) *Header {
    header := Header{}

    data := readNextBytes(file, int(unsafe.Sizeof(header)))
    buffer := bytes.NewBuffer(data)
    err := binary.Read(buffer, binary.BigEndian, &header)
    if err != nil {
        fmt.Println("binary.Read failed", err)
    }

    return &header
}

func showHeader(header *Header) {
    fmt.Printf("Constant : %X\n", header.Constant)
    fmt.Printf("PRG ROM SIZE : %x\n", header.PRG_ROM_SIZE)
    fmt.Printf("CHR ROM SIZE : %x\n", header.CHR_ROM_SIZE)
    fmt.Printf("Flags6 : %08b\n", header.Flags6)
    fmt.Printf("Flags7 : %08b\n", header.Flags7)
    fmt.Printf("PRG RAM SIZE : %x\n", header.PRG_RAM_SIZE)
    fmt.Printf("Flags9 : %08b\n", header.Flags9)
    fmt.Printf("Flags10 : %08b\n\n", header.Flags10)
}

func readNextBytes(file *os.File, number int) []byte {
    bytes := make([]byte, number)

    _, err := file.Read(bytes)
    if err != nil {
        fmt.Println("Reading Error")
        return nil
    }
    return bytes
}

func main(){
    argc := len(os.Args)
    if argc < 2 {
        fmt.Println("input file name.")
        return
    }

    filename := os.Args[1]
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("Err : ", err)
        return
    }
    defer f.Close()

    header := readHeader(f)
    showHeader(header)

    programROM := readNextBytes(f, int(header.PRG_ROM_SIZE) * 0x4000)
    characterROM := readNextBytes(f, int(header.CHR_ROM_SIZE) * 0x2000)
    fmt.Printf("Program ROM size : %x\n", len(programROM))
    fmt.Printf("Chatacter ROM size : %x\n", len(characterROM))

    fmt.Printf("Program ROM : %x\n### END OF PROGRAM ROM ###\n", programROM)
    fmt.Printf("Character ROM : %x\n### END OF CHARACTER ROM ###\n", characterROM)

}

