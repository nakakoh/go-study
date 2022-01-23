package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Movie struct {
		Title string
		// フィールドタグ `json:**`
		Year int `json:"released"` // JSON変換時にフィールド名を released に変える
		// omitempty はフィールドがゼロ値(boolはfalse)か空であればJSON出力しない
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Hmmphery Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	// 構造体をJSONに変換　マーシャリング (marshaling)
	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	}
	// 人が読みやすい形にマーシャリング
	{
		data, err := json.MarshalIndent(movies, "", "\t")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		// アンマーシャリング (Unmarshaling)
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
	}
}
