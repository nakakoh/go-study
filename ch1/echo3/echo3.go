package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	//fmt.Println(strings.Join(os.Args[1:], " "))
}
