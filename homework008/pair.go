package homework008

import (
	"fmt"
	"sort"
)

func Task2() {
	fmt.Println("homework008_2")
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}

	number := []int{}
	alphabet := []string{}
	for k, v := range m {
		if k != 3 {
			number = append(number, k)
			alphabet = append(alphabet, v)
		}
		sort.Ints(number)
		sort.Strings(alphabet)
	}
	fmt.Println(number)
	fmt.Println(alphabet)
}
