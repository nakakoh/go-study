package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // x == 0 ならパニック
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

/*
実行するとStdoutは以下のようになる
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3

StdErrは以下のように出力される
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x10c1220)
        /defer1/main.go:10 +0x157
main.f(0x1)
        /defer1/main.go:12 +0x132
main.f(0x2)
        /defer1/main.go:12 +0x132
main.f(0x3)
        /defer1/main.go:12 +0x132
main.main()
        /defer1/main.go:6 +0x1e
exit status 2
*/
