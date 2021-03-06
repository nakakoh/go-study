package main

import (
	"bytes"
	"fmt"
)

// IIntSet は負ではない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint64
}

// Has は負ではない値 x をセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add はセットに負ではない値 x を追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith は、 s と t の和集合を s に設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String は"{1 2 3}"の形式の文字列としいてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9 42}

	x.UnionWith(&y)
	fmt.Println(x.String()) // {1 9 42 144}

	fmt.Println(x.Has(9), x.Has(123)) // true false

	fmt.Println(&x)         // {1 9 42 144}
	fmt.Println(x.String()) // {1 9 42 144}
	fmt.Println(x)          // {[4398046511618 0 65536]}
}
