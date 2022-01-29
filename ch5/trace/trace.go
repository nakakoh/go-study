package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // 最後の追加の丸括弧を忘れないこと
	// ...大量の処理...
	time.Sleep(10 * time.Second) // スリープによって遅い操作を模倣
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
