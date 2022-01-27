package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer は URL のサーバーへ接続を試みます。
// 指数バックオフ（エクスポネンシャルバックオフ）を使って一分間試みます。
// 全ての試みが失敗したらエラーを報告します。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数バックオフ
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// logパッケージでもよい (自動で改行付加)
		log.Printf("Site is down: %v", err)
	}
}
