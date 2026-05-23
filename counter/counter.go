package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i <= 15; i++ {
        // %04b means: print as binary, pad with zeros to 4 places
        fmt.Printf("Decimal: %2d | Binary Bits: [ %04b ]\n", i, i)
        time.Sleep(500 * time.Millisecond)
    }
}
