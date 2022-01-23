package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Parse(time.RFC3339, "2019-10-15T04:41:17.52438Z"))
	fmt.Println(time.Parse(time.RFC3339Nano, "2019-10-15T04:41:17.52438Z"))

	t, _ := time.Parse(time.RFC3339, "2019-10-15T04:41:17.52438Z")
	fmt.Println(t)
}
