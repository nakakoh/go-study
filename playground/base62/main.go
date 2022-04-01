package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// uuid から Base62[A-Za-z0-9] の文字列を生成する
func main() {
	//str := time.Now().String()
	//b := getSha256(str)
	id := uuid.New()
	b, _ := id.MarshalBinary()

	var i big.Int
	i.SetBytes(b)

	fmt.Println("str:", id.String())
	fmt.Println("Base16:", i.Text(16))
	fmt.Println("Base62:", i.Text(62))
}

func getSha256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
