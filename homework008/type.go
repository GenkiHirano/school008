package homework008

import (
	"fmt"
)

func typeJudgment(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Printf("%vはそれ以外です。\n", x)
	case int, uint:
		fmt.Printf("%vは数値型です。\n", x)
	case string:
		fmt.Printf("%vは文字列型です。\n", x)
	}
}

func Task3() {
	fmt.Println("homework008_3")
	typeJudgment(100)
	typeJudgment("Hello, World!")
	typeJudgment(true)
}
