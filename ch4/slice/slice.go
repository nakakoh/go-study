package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August", 9: "September",
		10: "October", 11: "Novenber", 12: "December"}
	fmt.Println(months)
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Printf("data: %v, len: %d, cap: %d\n", months, len(months), cap(months))
	fmt.Printf("data: %v, len: %d, cap: %d\n", Q2, len(Q2), cap(Q2))
	fmt.Printf("data: %v, len: %d, cap: %d\n", summer, len(summer), cap(summer))

	// 共通要素を求める（非効率な）検査
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//fmt.Println(summer[:20])    // パニック: 範囲外
	endlessSummer := summer[:5] // （容量の範囲内で）スライスを拡張
	fmt.Println(endlessSummer)  // [June July August September October]

	// スライスは == で比較できない
	//fmt.Println(months[3:5] == months[3:5]) // コンパイルエラー

	// バイトスライス []byte は bytes.Equal関数で比較できる
	// それ以外は自前で実装が必要
	fmt.Println(equal(months[3:5], months[3:5])) // true

	// nilとの比較はできる
	// スライスが空かどうかは s == nil ではなく len(s) == 0 を使う
	{
		var s []int // len(s) == 0, s == nil
		fmt.Printf("len(s) = %d, s == nil: %t\n", len(s), s == nil)
		s = nil // len(s) == 0, s == nil
		fmt.Printf("len(s) = %d, s == nil: %t\n", len(s), s == nil)
		s = []int(nil) // len(s) == 0, s == nil
		fmt.Printf("len(s) = %d, s == nil: %t\n", len(s), s == nil)
		s = []int{} // len(s) == 0, s != nil
		fmt.Printf("len(s) = %d, s == nil: %t\n", len(s), s == nil)
	}

	// makeでスライス作成
	sl1 := make([]string, 5)
	sl2 := make([]string, 5, 10)
	fmt.Printf("len(s) = %d, s == nil: %t\n", len(sl1), sl1 == nil)
	fmt.Printf("len(s) = %d, s == nil: %t\n", len(sl2), sl2 == nil)

}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
