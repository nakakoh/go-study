package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func diff(a, b *[32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		tmp := uint(a[i] ^ b[i]) // XOR
		// 標準パッケージで楽をした。。
		sum += bits.OnesCount(tmp)
	}
	return sum
}

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", a, b)
	fmt.Printf("diff count: %d\n", diff(&a, &b))
}
