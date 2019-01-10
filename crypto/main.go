package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("Hello World\n"))
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, 1)
	s2 := hex.EncodeToString(b)
	fmt.Println(b)
	fmt.Println(s2)
}
