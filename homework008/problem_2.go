package homework008

import (
	"fmt"
)

func task2() {
	m := map[int]string{1: "Go", 2: "Ruby", 3: "JavaScript"}
	fmt.Println(m)

	// mから「3:JavaScript」を削除し、出力してください。

}

func task3() {

	names := []string{"佐藤", "田中", "鈴木"}
	fmt.Println(names)

	// namesに「園田」と「徳島」を追加してください。

	// namesを全件出力してください。

	// namesの要素数を出力してください。

}

func Test2() {
	fmt.Println("homework008_2")
	task2()
	task3()
}
