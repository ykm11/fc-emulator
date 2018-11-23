package testdata

import (
    "testing"
    "os"

    "../processor"
)

func TestHeader(t *testing.T) {
    f, err := os.Open("../../sample1.nes")
    if err != nil {
        t.Error("File open error")
    }
    defer f.Close()
    header := processor.ReadHeader(f)

    if header.Constant[0] != 0x4E || header.Constant[1] != 0x45 || header.Constant[2] != 0x53 || header.Constant[3] != 0x1A {
        t.Fatal("Wrong")
    }
}
