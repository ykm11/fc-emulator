package main

import (
    "fmt"
    "os"

    "./processor"
)


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

    // Prorogue
    header := processor.ReadHeader(f)
    header.ShowHeader()

    programROM := processor.ReadNextBytes(f, int(header.PRG_ROM_SIZE) * 0x4000)
    characterROM := processor.ReadNextBytes(f, int(header.CHR_ROM_SIZE) * 0x2000)
    fmt.Printf("Program ROM size : 0x%x\n", len(programROM))
    fmt.Printf("Chatacter ROM size : 0x%x\n", len(characterROM))


    //fmt.Printf("Program ROM : %x\n### END OF PROGRAM ROM ###\n", programROM)
    //fmt.Printf("Character ROM : %x\n### END OF CHARACTER ROM ###\n", characterROM)

    cpu := processor.CPU{
    }
    cpu.RESET()
    cpu.ShowRegister()

    cpu.MemoryMapping(programROM, 0x8000)
    fmt.Printf("%X\n", cpu.Memory[0x8000 : 0x8080])

    cpu.Run()

}

