package validator

import (
	"regexp"
	"testing"

	"github.com/go-playground/validator/v10"
)

/* Benchmark 結果
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz

Benchmark_Alphanum-8                                             2637794               403.1 ns/op            16 B/op          1 allocs/op
Benchmark_AlphanumMin6Max16-8                                    2197718               513.4 ns/op            16 B/op          1 allocs/op
Benchmark_AlphanumericMin6Max16ContainsAnyNumAndAlphabet-8       1967143               622.3 ns/op            16 B/op          1 allocs/op
Benchmark_Regex-8                                                4750104               274.6 ns/op             0 B/op          0 allocs/op
Benchmark_CompareRune-8                                         35721045                39.79 ns/op            0 B/op          0 allocs/op
Benchmark_CustomRegexValidator-8                                 1657042               776.1 ns/op            16 B/op          1 allocs/op
Benchmark_CustomRuneValidator-8                                  2115559               576.5 ns/op            16 B/op          1 allocs/op
*/

// 英数字チェック
// （英数混合必須チェックしていない）
func Benchmark_Alphanum(b *testing.B) {
	validate := validator.New()
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "required,alphanum")
	}
}

// 英数字と文字数チェック
// （英数混合必須チェックしていない）
func Benchmark_AlphanumMin6Max16(b *testing.B) {
	validate := validator.New()
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "required,alphanum,min=6,max=16")
	}
}

// 英数字と文字数、英数混合必須チェック
// containsany でいずれかが含まれているかチェックしているが、表記が長い。。
// containsany の中身をミスしそう...
func Benchmark_AlphanumericMin6Max16ContainsAnyNumAndAlphabet(b *testing.B) {
	validate := validator.New()
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "required,alphanum,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,min=6,max=16")
	}
}

// 正規表現による英数字と文字数チェック
// （英数混合必須チェックしていない）
// (go だと肯定的先読みサポートがないので正規表現だけで一発ではできないらしい...)
func Benchmark_Regex(b *testing.B) {
	r := regexp.MustCompile(`^[0-9a-zA-Z]{6,16}$`)
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = r.MatchString(s)
	}
}

// Runeによる英数字と文字数、英数混合必須チェック
// ロジック分かりづらいが最速
func Benchmark_CompareRune(b *testing.B) {
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var numCount, alphabetCount int
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				alphabetCount++
			} else if c >= 'a' && c <= 'z' {
				alphabetCount++
			} else if c >= '0' && c <= '9' {
				numCount++
			} else {
				// error
				panic("error!")
			}
		}
		if numCount == 0 || alphabetCount == 0 || len(s) < 6 || len(s) > 16 {
			// error
			panic("error!")
		}
	}
}

// 正規表現によるカスタムバリデーターによる英数字と文字数、英数混合必須チェック
// 事前コンパイルしておけば、そんなに遅くなり
func Benchmark_CustomRegexValidator(b *testing.B) {
	validate := validator.New()
	validate.RegisterValidation("include_alphabet", includeAlphabet)
	validate.RegisterValidation("include_number", includeNumber)
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "required,alphanum,min=6,max=16,include_alphabet,include_number")
	}
}

var alphaRegex = regexp.MustCompile("[a-zA-Z]")

func includeAlphabet(fl validator.FieldLevel) bool {
	return alphaRegex.MatchString(fl.Field().String())
}

var numberRegex = regexp.MustCompile("[0-9]")

func includeNumber(fl validator.FieldLevel) bool {
	return numberRegex.MatchString(fl.Field().String())
}

// rune によるカスタムバリデーターによる英数字と文字数、英数混合必須チェック
// 若干ロジック複雑（といっても単純）だが、ロジックは使う側は直接見ないし
// 標準の validator と差があまりないのでよさそう
func Benchmark_CustomRuneValidator(b *testing.B) {
	validate := validator.New()
	validate.RegisterValidation("include_alphabet", includeAlphabetRune)
	validate.RegisterValidation("include_number", includeNumberRune)
	s := "abcde678901ABCDE"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = validate.Var(&s, "required,alphanum,min=6,max=16,include_alphabet,include_number")
	}
}

func includeAlphabetRune(fl validator.FieldLevel) bool {
	for _, c := range fl.Field().String() {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			return true
		}
	}
	return false
}

func includeNumberRune(fl validator.FieldLevel) bool {
	for _, c := range fl.Field().String() {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
