package processor

type Instruction struct {
    mode string
    syntax string
    clk uint8
}

// ### Addressing Mode ###
//
// Accumulator Addressing           -> accumulator
// Immediate Addressing             -> immediate
// Absolute Addressing              -> absolute
// Zero Page Addressing             -> zeroPage
// Indexed Zero Page Addressing X   -> zeroPageX
// Indexed Zero Page Addressing Y   -> zeroPageY
// Indexed Absolute Addressing X    -> absoluteX
// Indexed Absolute Addressing Y    -> absoluteY
// Implied Addressing               -> implied
// Relative Addressing              -> relative
// Indexed Indirect Addressing      -> indexedIndirectX
// Indirect Indexed Addressing      -> indirectIndexedY
// Absolute Indirect Addressing     -> absoluteIndirect

func GetInstructions() map[uint8]Instruction {
    codes := make(map[uint8]Instruction)

    // http://pgate1.at-ninja.jp/NES_on_FPGA/nes_cpu.htm#clock

    // ADC
    codes[0x69] = Instruction{ mode:"immediate", syntax:"ADC", clk:2 }
    codes[0x65] = Instruction{ mode:"zeroPage", syntax:"ADC", clk:3 }
    codes[0x75] = Instruction{ mode:"zeroPageX", syntax:"ADC", clk:4 }
    codes[0x6D] = Instruction{ mode:"absolute", syntax:"ADC", clk:4 }
    codes[0x7D] = Instruction{ mode:"absoluteX", syntax:"ADC", clk:5 }
    codes[0x79] = Instruction{ mode:"absoluteY", syntax:"ADC", clk:5 }
    codes[0x61] = Instruction{ mode:"indexedIndirectX", syntax:"ADC", clk:6 }
    codes[0x71] = Instruction{ mode:"indirectIndexedY", syntax:"ADC", clk:6 }

    // SBC
    codes[0xE9] = Instruction{ mode:"immediate", syntax:"SBC", clk:2 }
    codes[0xE5] = Instruction{ mode:"zeroPage", syntax:"SBC", clk:3 }
    codes[0xF5] = Instruction{ mode:"zeroPageX", syntax:"SBC", clk:4 }
    codes[0xED] = Instruction{ mode:"absolute", syntax:"SBC", clk:4 }
    codes[0xFD] = Instruction{ mode:"absoluteX", syntax:"SBC", clk:5 }
    codes[0xF9] = Instruction{ mode:"absoluteY", syntax:"SBC", clk:5 }
    codes[0xE1] = Instruction{ mode:"indexedIndirectX", syntax:"SBC", clk:6 }
    codes[0xF1] = Instruction{ mode:"indirectIndexedY", syntax:"SBC", clk:6 }

    // AND
    codes[0x29] = Instruction{ mode:"immediate", syntax:"AND", clk:2 }
    codes[0x25] = Instruction{ mode:"zeroPage", syntax:"AND", clk:3 }
    codes[0x35] = Instruction{ mode:"zeroPageX", syntax:"AND", clk:4 }
    codes[0x2D] = Instruction{ mode:"absolute", syntax:"AND", clk:4 }
    codes[0x3D] = Instruction{ mode:"absoluteX", syntax:"AND", clk:5 }
    codes[0x39] = Instruction{ mode:"absoluteY", syntax:"AND", clk:5 }
    codes[0x21] = Instruction{ mode:"indexedIndirectX", syntax:"AND", clk:6 }
    codes[0x31] = Instruction{ mode:"indirectIndexedY", syntax:"AND", clk:6 }

    // ORA
    codes[0x09] = Instruction{ mode:"immediate", syntax:"ORA", clk:2 }
    codes[0x05] = Instruction{ mode:"zeroPage", syntax:"ORA", clk:3 }
    codes[0x15] = Instruction{ mode:"zeroPageX", syntax:"ORA", clk:4 }
    codes[0x0D] = Instruction{ mode:"absolute", syntax:"ORA", clk:4 }
    codes[0x1D] = Instruction{ mode:"absoluteX", syntax:"ORA", clk:5 }
    codes[0x19] = Instruction{ mode:"absoluteY", syntax:"ORA", clk:5 }
    codes[0x01] = Instruction{ mode:"indexedIndirectX", syntax:"ORA", clk:6 }
    codes[0x11] = Instruction{ mode:"indirectIndexedY", syntax:"ORA", clk:6 }

    // EOR
    codes[0x49] = Instruction{ mode:"immediate", syntax:"EOR", clk:2 }
    codes[0x45] = Instruction{ mode:"zeroPage", syntax:"EOR", clk:3 }
    codes[0x55] = Instruction{ mode:"zeroPageX", syntax:"EOR", clk:4 }
    codes[0x4D] = Instruction{ mode:"absolute", syntax:"EOR", clk:4 }
    codes[0x5D] = Instruction{ mode:"absoluteX", syntax:"EOR", clk:5 }
    codes[0x59] = Instruction{ mode:"absoluteY", syntax:"EOR", clk:5 }
    codes[0x41] = Instruction{ mode:"indexedIndirectX", syntax:"EOR", clk:6 }
    codes[0x51] = Instruction{ mode:"indirectIndexedY", syntax:"EOR", clk:6 }

    // ASL
    codes[0x0A] = Instruction{ mode:"accumulator", syntax:"ASL", clk:2 }
    codes[0x06] = Instruction{ mode:"zeroPage", syntax:"ASL", clk:5 }
    codes[0x16] = Instruction{ mode:"zeroPageX", syntax:"ASL", clk:6 }
    codes[0x0E] = Instruction{ mode:"absolute", syntax:"ASL", clk:6 }
    codes[0x1E] = Instruction{ mode:"absoluteX", syntax:"ASL", clk:7 }

    // LSR
    codes[0x4A] = Instruction{ mode:"accumulator", syntax:"LSR", clk:2 }
    codes[0x46] = Instruction{ mode:"zeroPage", syntax:"LSR", clk:5 }
    codes[0x56] = Instruction{ mode:"zeroPageX", syntax:"LSR", clk:6 }
    codes[0x4E] = Instruction{ mode:"absolute", syntax:"LSR", clk:6 }
    codes[0x5E] = Instruction{ mode:"absoluteX", syntax:"LSR", clk:7 }

    // ROL
    codes[0x2A] = Instruction{ mode:"accumulator", syntax:"ROL", clk:2 }
    codes[0x26] = Instruction{ mode:"zeroPage", syntax:"ROL", clk:5 }
    codes[0x36] = Instruction{ mode:"zeroPageX", syntax:"ROL", clk:6 }
    codes[0x2E] = Instruction{ mode:"absolute", syntax:"ROL", clk:6 }
    codes[0x3E] = Instruction{ mode:"absoluteX", syntax:"ROL", clk:7 }

    // ROR
    codes[0x6A] = Instruction{ mode:"accumulator", syntax:"ROR", clk:2 }
    codes[0x66] = Instruction{ mode:"zeroPage", syntax:"ROR", clk:5 }
    codes[0x76] = Instruction{ mode:"zeroPageX", syntax:"ROR", clk:6 }
    codes[0x6E] = Instruction{ mode:"absolute", syntax:"ROR", clk:6 }
    codes[0x7E] = Instruction{ mode:"absoluteX", syntax:"ROR", clk:7 }

    // for Relative
    codes[0x90] = Instruction{ mode:"relative", syntax:"BCC", clk:3 }
    codes[0xB0] = Instruction{ mode:"relative", syntax:"BCS", clk:3 }
    codes[0xF0] = Instruction{ mode:"relative", syntax:"BEQ", clk:3 }
    codes[0xD0] = Instruction{ mode:"relative", syntax:"BNE", clk:3 }
    codes[0x50] = Instruction{ mode:"relative", syntax:"BVC", clk:3 }
    codes[0x70] = Instruction{ mode:"relative", syntax:"BVS", clk:3 }
    codes[0x10] = Instruction{ mode:"relative", syntax:"BPL", clk:3 }
    codes[0x30] = Instruction{ mode:"relative", syntax:"BMI", clk:3 }

    // BIT
    codes[0x24] = Instruction{ mode:"zeroPage" , syntax:"BIT", clk:3 }
    codes[0x2C] = Instruction{ mode:"absolute" , syntax:"BIT", clk:3 }

    // JMP
    codes[0x4C] = Instruction{ mode:"absolute", syntax:"JMP", clk:3 }
    codes[0x6C] = Instruction{ mode:"absoluteIndirect", syntax:"JMP", clk:5 }

    // JSR
    codes[0x20] = Instruction{ mode:"absolute", syntax:"JSR", clk:6 }

    // RTS
    codes[0x60] = Instruction{ mode:"implied", syntax:"RTS", clk:6 }

    // for Implied
    codes[0x00] = Instruction{ mode:"implied", syntax:"BRK", clk:7 }
    codes[0x40] = Instruction{ mode:"implied", syntax:"RTI", clk:6 }

    // CMP
    codes[0xC9] = Instruction{ mode:"immediate", syntax:"CMP", clk:2 }
    codes[0xC5] = Instruction{ mode:"zeroPage", syntax:"CMP", clk:3 }
    codes[0xD5] = Instruction{ mode:"zeroPageX", syntax:"CMP", clk:4 }
    codes[0xCD] = Instruction{ mode:"absolute", syntax:"CMP", clk:4 }
    codes[0xDD] = Instruction{ mode:"absoluteX", syntax:"CMP", clk:5 }
    codes[0xD9] = Instruction{ mode:"absoluteY", syntax:"CMP", clk:5 }
    codes[0xC1] = Instruction{ mode:"indexedIndirectX", syntax:"CMP", clk:6 }
    codes[0xD1] = Instruction{ mode:"indirectIndexedY", syntax:"CMP", clk:6 }

    // CPX
    codes[0xE0] = Instruction{ mode:"immediate", syntax:"CPX", clk:2 }
    codes[0xE4] = Instruction{ mode:"zeroPage", syntax:"CPX", clk:3 }
    codes[0xEC] = Instruction{ mode:"absolute", syntax:"CPX", clk:4 }

    // CPY
    codes[0xC0] = Instruction{ mode:"immediate", syntax:"CPY", clk:2 }
    codes[0xC4] = Instruction{ mode:"zeroPage", syntax:"CPY", clk:3 }
    codes[0xCC] = Instruction{ mode:"absolute", syntax:"CPY", clk:4 }

    // INC
    codes[0xE6] = Instruction{ mode:"zeroPage", syntax:"INC", clk:5 }
    codes[0xF6] = Instruction{ mode:"zeroPageX", syntax:"INC", clk:6 }
    codes[0xEE] = Instruction{ mode:"absolute", syntax:"INC", clk:6 }
    codes[0xFE] = Instruction{ mode:"absoluteX", syntax:"INC", clk:7 }

    // DEC
    codes[0xC6] = Instruction{ mode:"zeroPage", syntax:"DEC", clk:5 }
    codes[0xD6] = Instruction{ mode:"zeroPageX", syntax:"DEC", clk:6 }
    codes[0xCE] = Instruction{ mode:"absolute", syntax:"DEC", clk:6 }
    codes[0xDE] = Instruction{ mode:"absoluteX", syntax:"DEC", clk:7 }

    // for Implied
    codes[0xE8] = Instruction{ mode:"immediate", syntax:"INX", clk:2 }
    codes[0xCA] = Instruction{ mode:"immediate", syntax:"DEX", clk:2 }
    codes[0xC8] = Instruction{ mode:"immediate", syntax:"INY", clk:2 }
    codes[0x88] = Instruction{ mode:"immediate", syntax:"DEY", clk:2 }

    // for Implied
    codes[0x18] = Instruction{ mode:"immediate", syntax:"CLC", clk:2 }
    codes[0x38] = Instruction{ mode:"immediate", syntax:"SEC", clk:2 }
    codes[0x58] = Instruction{ mode:"immediate", syntax:"CLI", clk:2 }
    codes[0x78] = Instruction{ mode:"immediate", syntax:"SEI", clk:2 }
    codes[0xD8] = Instruction{ mode:"immediate", syntax:"CLD", clk:2 }
    codes[0xF8] = Instruction{ mode:"immediate", syntax:"SED", clk:2 }
    codes[0xB8] = Instruction{ mode:"immediate", syntax:"CLV", clk:2 }

    // LDA
    codes[0xA9] = Instruction{ mode:"immediate", syntax:"LDA", clk:2 }
    codes[0xA5] = Instruction{ mode:"zeroPage", syntax:"LDA", clk:3 }
    codes[0xB5] = Instruction{ mode:"zeroPageX", syntax:"LDA", clk:4 }
    codes[0xAD] = Instruction{ mode:"absolute", syntax:"LDA", clk:4 }
    codes[0xBD] = Instruction{ mode:"absoluteX", syntax:"LDA", clk:5 }
    codes[0xB9] = Instruction{ mode:"absoluteY", syntax:"LDA", clk:5 }
    codes[0xA1] = Instruction{ mode:"indexedIndirectX", syntax:"LDA", clk:6 }
    codes[0xB1] = Instruction{ mode:"indirectIndexedY", syntax:"LDA", clk:6 }

    // LDX
    codes[0xA2] = Instruction{ mode:"immediate", syntax:"LDX", clk:2 }
    codes[0xA6] = Instruction{ mode:"zeroPage", syntax:"LDX", clk:3 }
    codes[0xB6] = Instruction{ mode:"zeroPageY", syntax:"LDX", clk:4 }
    codes[0xAE] = Instruction{ mode:"absolute", syntax:"LDX", clk:4 }
    codes[0xBE] = Instruction{ mode:"absoluteY", syntax:"LDX", clk:5 }

    // LDY
    codes[0xA0] = Instruction{ mode:"immediate", syntax:"LDY", clk:2 }
    codes[0xA4] = Instruction{ mode:"zeroPage", syntax:"LDY", clk:3 }
    codes[0xB4] = Instruction{ mode:"zeroPageX", syntax:"LDY", clk:4 }
    codes[0xAC] = Instruction{ mode:"absolute", syntax:"LDY", clk:4 }
    codes[0xBC] = Instruction{ mode:"absoluteX", syntax:"LDY", clk:5 }

    // STA
    codes[0x85] = Instruction{ mode:"zeroPage", syntax:"STA", clk:3 }
    codes[0x95] = Instruction{ mode:"zeroPageX", syntax:"STA", clk:4 }
    codes[0x8D] = Instruction{ mode:"absolute", syntax:"STA", clk:4 }
    codes[0x9D] = Instruction{ mode:"absoluteX", syntax:"STA", clk:5 }
    codes[0x99] = Instruction{ mode:"absoluteY", syntax:"STA", clk:5 }
    codes[0x81] = Instruction{ mode:"indexedIndirectX", syntax:"STA", clk:6 }
    codes[0x91] = Instruction{ mode:"indirectIndexedY", syntax:"STA", clk:6 }

    // STX
    codes[0x86] = Instruction{ mode:"zeroPage", syntax:"STX", clk:3 }
    codes[0x96] = Instruction{ mode:"zeroPageY", syntax:"STX", clk:4 }
    codes[0x8E] = Instruction{ mode:"absolute", syntax:"STX", clk:4 }

    // STY
    codes[0x84] = Instruction{ mode:"zeroPage", syntax:"STY", clk:3 }
    codes[0x94] = Instruction{ mode:"zeroPageY", syntax:"STY", clk:4 }
    codes[0x8C] = Instruction{ mode:"absolute", syntax:"STY", clk:4 }

    // for Implied
    codes[0xAA] = Instruction{ mode:"implied", syntax:"TAX", clk:2 }
    codes[0x8A] = Instruction{ mode:"implied", syntax:"TXA", clk:2 }
    codes[0xA8] = Instruction{ mode:"implied", syntax:"TAY", clk:2 }
    codes[0x98] = Instruction{ mode:"implied", syntax:"TYA", clk:2 }
    codes[0x9A] = Instruction{ mode:"implied", syntax:"TXS", clk:2 }
    codes[0xBA] = Instruction{ mode:"implied", syntax:"TSX", clk:2 }

    codes[0x48] = Instruction{ mode:"implied", syntax:"PHA", clk:3 }
    codes[0x68] = Instruction{ mode:"implied", syntax:"PLA", clk:4 }
    codes[0x08] = Instruction{ mode:"implied", syntax:"PHP", clk:3 }
    codes[0x28] = Instruction{ mode:"implied", syntax:"PLP", clk:4 }

    codes[0xEA] = Instruction{ mode:"implied", syntax:"NOP", clk:2 }


    return codes
}

