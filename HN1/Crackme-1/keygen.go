package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    // Define the array to hold the individual key bytes
    key := make([]byte, 10)

    // seed current system time
    rand.Seed(time.Now().Unix())

    // generate (some) random bytes meeting the required range conditions
    key[0] = byte(rand.Intn(122 - 0x48)) + 0x48
    key[1] = byte(rand.Intn(0x6c - 33)) + 33
    key[2] = 0x56
    key[3] = byte(rand.Intn(122 - 0x66)) + 0x66
    key[4] = byte(rand.Intn(0x33 - 33)) + 33
    key[5] = 0x7a
    key[6] = byte(rand.Intn(122 - 0x38)) + 0x38
    key[7] = byte(rand.Intn(0x4d - 33)) + 33
    key[8] = byte(rand.Intn(122 - 0x53)) + 0x53
    key[9] = 0x32

    fmt.Println("Key:", string(key))
}
