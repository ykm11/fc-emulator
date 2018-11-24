package testdata

import (
    "testing"
)

func TestCasting(t *testing.T) {

    if uint8(0x1234 & 0xFF) != 0x34 {
        t.Fatal("Cating Error")
    }

}
