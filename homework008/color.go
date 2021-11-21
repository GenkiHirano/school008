package homework008

import (
	"fmt"
)

func likeColor(s string) {
	if s == "佐藤" {
		fmt.Println("佐藤さんは赤色が好き")
	} else if s == "鈴木" {
		fmt.Println("鈴木さんは青色が好き")
	} else if s == "田中" {
		fmt.Println("田中さんの好きな色は知りません")
	}
}

func Task1() {
	fmt.Println("homework008_1")
	students := [3]string{"佐藤", "鈴木", "田中"}
	for _, v := range students {
		likeColor(v)
	}
}
