package main

import "fmt"

func main() {
	fmt.Printf("%d\n", gcd(1, 2))
	fmt.Printf("%d\n", gcd(20, 150))
	fmt.Printf("%d\n", gcd(34, 56))
	fmt.Printf("%d\n", fib(1))
	fmt.Printf("%d\n", fib(10))
	fmt.Printf("%d\n", fib(33))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
