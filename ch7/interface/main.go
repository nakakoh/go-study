package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {

	{
		var w io.Writer
		w = os.Stdout         // OK: *os.File は Write メソッドあり
		w = new(bytes.Buffer) // OK: *bytes.Buffer は Write めそっdおあり
		w = time.Second       // コンパイルエラー: time.Duration は Write メソッドなし

		var rwc io.ReadWriteCloser
		rwc = os.Stdout         // OK: *os.File は Read, Write, Close メソッドあり
		rwc = new(bytes.Buffer) // コンパイルエラー: *bytes.Buffer は Close なし

		w = rwc // OK: io.ReadWriteCloser は Write メソッドあり
		rwc = w // コンパイルエラー: io.Writer は Close メソッドなし
		// ReadWriter と ReadWriteCloser は Writer のすべてのメソッドを含んでいるので、
		// ReadWriter か ReadWriteCloser を満足する全ての型は必然的に Writer を満足する
	}

}

// 具象型 concrete types
// インタフェース型(interface type) は抽象型(abstract type)

// インタフェース
// fnt Fprintfの例
/*
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int error)
func Printf(format string, args ...interface{}) (int , error) {
	return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) (int , error) {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}

// Fprintf の最初のパラメータはio.Writer インタフェース
pakcage io
// Writer は、基本的なWriteメソッドを包んでいるインタフェース
type Writer interface {
	// Write は p から len(p) バイトを基底のデータストリームへ書き込みます。
	// p から書き込まれたバイト数(0 <= n <= len(p)) と、書き込みを早く終わらせた原因と
  // なったエラーを返します。
	// Write は、 n < len(p) であるような n を返す場合には nil でない error を返さなければなりません
	// Write は、たとえ一時的であってもスライスのデータを変更してはいけません。

  // 実装は、p を持ち続けてはいけません。
	Write(p []byte) (n int, err error)
}
*/

// 型に対してStringメソッドを宣言することで、最も広く使われているインタフェースの一つである
// fmt.Stringer を満足させることになる
/*
package fmt
type Stringer interface {
	String() string
}
*/
