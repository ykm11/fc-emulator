package processor

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

func ReadNextBytes(file *os.File, number int) []byte {
    bytes := make([]byte, number)

    _, err := file.Read(bytes)
    if err != nil {
        fmt.Println("Reading Error")
        return nil
    }
    return bytes
}

func ReadHeader(file *os.File) *Header {
    header := Header{}

    data := ReadNextBytes(file, int(unsafe.Sizeof(header)))
    buffer := bytes.NewBuffer(data)
    err := binary.Read(buffer, binary.BigEndian, &header)
    if err != nil {
        fmt.Println("binary.Read failed", err)
    }

    return &header
}

func (header *Header) ShowHeader() {
    fmt.Println("---------Header---------")
    fmt.Printf("Constant : %X\n", header.Constant)
    fmt.Printf("PRG ROM SIZE : %x\n", header.PRG_ROM_SIZE)
    fmt.Printf("CHR ROM SIZE : %x\n", header.CHR_ROM_SIZE)
    fmt.Printf("Flags6 : %08b\n", header.Flags6)
    fmt.Printf("Flags7 : %08b\n", header.Flags7)
    fmt.Printf("PRG RAM SIZE : %x\n", header.PRG_RAM_SIZE)
    fmt.Printf("Flags9 : %08b\n", header.Flags9)
    fmt.Printf("Flags10 : %08b\n", header.Flags10)
    fmt.Println("---------Header---------")
}
