package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var outputSha256 = flag.Bool("sha256", true, "convert stdin to SHA256")
var outputSha384 = flag.Bool("sha384", false, "convert stdin to SHA384")
var outputSha512 = flag.Bool("sha512", false, "convert stdin to SHA512")

func main() {
	flag.Parse()
	fmt.Print("input var: ")
	input := bufio.NewScanner(os.Stdin)
	var tmp string
	if input.Scan() {
		tmp = input.Text()
	}
	if *outputSha256 {
		fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(tmp)))
	}
	if *outputSha384 {
		fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(tmp)))
	}
	if *outputSha512 {
		fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(tmp)))
	}
}
