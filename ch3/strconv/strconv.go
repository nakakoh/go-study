package main

import (
	"fmt"
	"strconv"
)

func main() {
	{
		x := 123
		y := fmt.Sprintf("%d", x)
		fmt.Println(y, strconv.Itoa(x))             // 123 123
		fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011

		s := fmt.Sprintf("x=%b", x)
		fmt.Println(s)
	}

	x, _ := strconv.Atoi("123")
	y, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(x, y)
}
