package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

// sha256 の ハッシュ値 から Base62[A-Za-z0-9] の文字列を生成する
func main() {
	str := time.Now().String()
	b := getSha256(str)

	var i big.Int
	i.SetBytes(b)

	fmt.Println("str:", str)
	fmt.Println("Base16:", i.Text(16))
	fmt.Println("Base62:", i.Text(62))
}

func getSha256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
