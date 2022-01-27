package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 関数宣言の基本
// fun name(parameter-list) (result-list) {
//   body
// }
// go にはデフォルトパラメータ値という概念はない

func main() {
	fmt.Println(hypot(3, 4)) // 5

	// int型の呼び出し方パターン
	fmt.Printf("%T\n", add)   // func(int, int) int
	fmt.Printf("%T\n", sub)   // func(int, int) int
	fmt.Printf("%T\n", first) // func(int, int) int
	fmt.Printf("%T\n", zero)  // func(int, int) int

	// goは他の言語のような例外（Exception）を使って報告する仕組みはない
	// 基本的に関数の戻り値でエラーを返す方法が主流
	/*
		// 原因が一つしかありえない場合は結果はブーリアンとして、たいてい ok と名付ける
		value, ok := cache.Lookup(key)
		if !ok {
			// cache[key] が存在しない
		}
	*/

	// エラーは全体的な失敗に対する根本問題からの明確な因果の連鎖を提供すべき
	// error1: error2: error3: error4: ....
	/*
		doc ,err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
		}
	*/

	// ioクラスにはファイルの終わりを示す io.EOF を提供する
	{
		f, err := os.Open("./dummy.txt")
		if err != nil {
			log.Printf("opening file failure: %s", err)
			os.Exit(1)
		}
		in := bufio.NewReader(f)
		for {
			_, _, err := in.ReadRune()
			if err == io.EOF {
				log.Printf("end of file reached")
				break // 読み込みを終了
			}
		}
	}

	// 関数はファーストクラス値(first-class value)
	// 関数値(function value)は他の値と同様に型を持ち、変数に代入したり、関数から返したりできる
	{
		f := square
		fmt.Println(f(3)) // 9
		f = negative
		fmt.Println(f(3))     // -3
		fmt.Printf("%T\n", f) // func(int) int
		//f = product // コンパイルエラー: func(int, int) int を func(int) int へ代入できない
	}

	// 関数値のゼロ値は nil
	{
		var f func(int) int
		//f(3) // パニック: nil の関数呼び出し
		if f != nil {
			f(3)
		}
	}

	// 関数値をつかって振る舞いをパラメータ化できる
	{
		var add1 = func(r rune) rune { return r + 1 }
		fmt.Println(strings.Map(add1, "HAL-9000")) // IBM.:111
		fmt.Println(strings.Map(add1, "VMS"))      // WNT
		fmt.Println(strings.Map(add1, "Admix"))    // Benjy
	}

	// 無名関数
	{
		strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	}
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 以下の関数宣言は同じ
//func f(i, j, k int, s, t string)                { /* ... */ }
//func f(i int, j int, k int, s string, t string) { /* ... */ }

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

// CountWordsAndImages は HTML ドキュメントに対する HTTP GET リクエストを url へ
// 行い、そのドキュメントないに含まれる単語と画像の数を返します。
// 名前付きの結果を持つ関数内では return 文は省略できる。空リターン (bare return)
// ただ空リターンは重複は減らせるが、コードの理解を容易にしてくれることはほとんどない
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// dummy
	words = 10
	images = 20
	return
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }