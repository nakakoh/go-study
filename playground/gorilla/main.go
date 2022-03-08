package main

import (
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	sesskey := securecookie.GenerateRandomKey(32)
	fmt.Printf("sesskey: %s\n", sesskey)
	str := string(sesskey[:])
	fmt.Printf("sesskey: %s\n", str)
}
