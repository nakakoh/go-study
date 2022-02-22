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

	// 可変個引数関数
	// 可変個引数関数の型は普通のスライスとは明確に異なる
	{
		var f = func(...int) {}
		var g = func([]int) {}
		fmt.Printf("%T\n", f) // func(...int)
		fmt.Printf("%T\n", g) // func([]int)

		// 文字列のフォーマット関数で良く使われている
		var errorf = func(linenum int, format string, args ...interface{}) {
			fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
			fmt.Fprintf(os.Stderr, format, args...)
			fmt.Fprintln(os.Stderr)
		}
		linenum, name := 12, "count"
		errorf(linenum, "undefined: %s", name) // Line 12: undefined: count
	}

	// 遅延関数呼び出し
	// defer で関数宣言することで、関数完了後に実施する処理を定義できる
	{
		var double = func(x int) (result int) {
			defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
			return x + x
		}
		_ = double(4) // 8

		// 呼び出し元に返す値を変えることもできる
		var triple = func(x int) (result int) {
			defer func() { result += x }()
			return double(x)
		}
		fmt.Println(triple(4)) // 12
	}

	// パニック
	// 発生すると、そのゴルーチン内の全ての遅延関数の呼び出しを行い、プログラムはログメッセージを表示してCrashする
	// 通常はerror値を使って適切にエラー処理されるべき
	{
		panic("panic!")

		// regexp.MustCompileという関数は以下のような実装になっている
		// これはパッケージレベルの変数を初期化するのを便利にする
		// (なんか関数を実行する前に初期化すべきもののエラー検知的なもの？)
		/*
			package regexp
			func Compile(expr string) (*Regexp, error) { /* ... *\/ }
			func MustCompile(expr string) * Regexp {
				re, err := Compile(expr)
				if err != nil {
					panic(err)
				}
				return re
			}
		*/
		//var httpSchemeRE = regexp.MustCompile(`^https?:`) // "http:" or "https:"
	}

}

// カプセル化(encapsulated) // 情報隠蔽(information hiding)
// Goは名前の可視性を制御する仕組みは大文字で始まるかどうかだけで決まる
// 構造体では単一のフィールドしか無くても、以下のように小文字で定義すれば隠せる
type IntSet struct {
	words []uint64
}

// 構造体のフィールドは隠して、実装だけを提供することで互換性を保ちやすくなる
type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

// ゲッター(getter) セッター(setter)
type Logger struct {
	flags  int
	prefix string
}

func (l *Logger) Flags() int
func (l *Logger) SetFlags(flags int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)

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
