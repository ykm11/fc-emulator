package utils

func Compare(acc, M uint8) int8 {
    // オペランドを符号付きで計算するのかどうか怪しいがとりあえず符号付き
    result := int16(int8(acc)) - int16(int8(M))
    if result < 0 {
        return -1
    } else if result > 0 {
        return 1
    }
    return 0
}
