package main

import (
	"fmt"
	"math"
)

func main() {
	// リテラルを省略すると、前のリテラルが反映される
	{
		const (
			a = 1
			b
			c = 2
			d
		)
		fmt.Println(a, b, c, d) // 1 1 2 2
	}

	// iota 定数生成器 (constant generator)
	// ゼロから始まり順番に個々の項目ごとに1増加させる(timeパッケージの例)
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	// 0 1 2 3 4 5 6
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 符号なし整数の最下位5ビットの各ビットを立てる
	// netパッケージの例
	type Flags uint
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	// 1 2 4 8 16
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
	// 1 10 100 1000 10000
	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	// 1024の累乗に名前をつける
	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776 (1 << 32 を超える)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // (1 << 64 を超える)
		YiB
	)
	// ZiB, YiBはオーバーフロー(int64を超えるため)
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB /*, ZiB, YiB*/)
	// ZiB, YiBは精度を保っているので以下は可能
	fmt.Println(YiB / ZiB) // 1024

	// math.PI は型を当てると精度が下がる
	var x float32 = math.Pi
	var y float64 = math.Pi
	fmt.Println(x, y, math.Pi)

	// 型付けなし定数
	// リテラルに対して型が決まる
	{
		var f float64 = 212
		g := (f - 32) * 5 / 9
		fmt.Printf("%v %T\n", g, g) // 100; (f -32) * 5 は float64
		g = 5 / 9 * (f - 32)
		fmt.Printf("%v %T\n", g, g) // 0; 5/9 は型付けなし整数で 0
		g = 5.0 / 9.0 * (f - 32)
		fmt.Printf("%v %T\n", g, g) // 100; 5.0/9.0 は型付けなし浮動小数点数
	}

	{
		var f float64 = 3 + 0i // 型付けなし複素数 -> float64
		fmt.Printf("%v %T\n", f, f)
		f = 2 // 型付けなし整数 -> float64
		fmt.Printf("%v %T\n", f, f)
		f = 1e123 // 型付けなし浮動小数点数 -> float64
		fmt.Printf("%v %T\n", f, f)
		f = 'a' // 型付けなしルーン -> float64
		fmt.Printf("%v %T\n", f, f)
	}

	{
		const (
			deafbeef = 0xdeadbeef        // 型付けなしintで値は3735928559
			a        = uint32(deafbeef)  // uint32 3735928559
			b        = float32(deafbeef) // float32 3.7359286e+09 (丸めあげ)
			c        = float64(deafbeef) // float64 3.735928559e+09 (正確)
			// 以降コンパイルエラー
			//d = int32(deafbeef) // 定数がint32をオーバーフロー
			//e = float64(1e309)  // 定数がfloat64をオーバーフロー
			//f = uint(-1)        // 定数がuintをアンダーフロー
		)
		fmt.Println(deafbeef, a, b, c)
	}

	{
		i := 0      // 型付けなし整数;         暗黙的に int(0)
		r := '\000' // 型付けなしルーン;       暗黙的に rune('\000')
		f := 0.0    // 型付けなし浮動小数点数; 暗黙的に float64(0.0)
		c := 0i     // 型付けなし複素数;       暗黙的に complex128(0i)
		fmt.Printf("%v %T\n", i, i)
		fmt.Printf("%v %T\n", r, r)
		fmt.Printf("%v %T\n", f, f)
		fmt.Printf("%v %T\n", c, c)
		// intはサイズが決まらないので注意
	}
}
