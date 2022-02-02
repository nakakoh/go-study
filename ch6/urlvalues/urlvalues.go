package main

import (
	"fmt"
	"net/url"
)

func main() {

	m := url.Values{"lang": {"en"}} // 直接の構築
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (最初の値)
	fmt.Println(m["item"])     // "[1 2]" (直接のマップ・アクセス)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // パニック: nilマップのエントリへの代入
}
