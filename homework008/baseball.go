package homework008

import (
	"fmt"
)

var st, bl int

func Task4() {
	var i int
	fmt.Println("homework008_4")
	pitching := [5]string{"strike", "ball", "ball", "strike", "strike"}

	for i = 0; i < 5; i++ {
		count(pitching[i])
	}
}

func count(x string) {
	if x == "strike" && st < 2 {
		fmt.Println("strike!")
		st++
	} else if x == "strike" && st == 2 {
		fmt.Println("strike!")
		fmt.Println("out!")
	} else if x == "ball" && bl < 3 {
		fmt.Println("ball!")
		bl++
	} else if x == "ball" && bl == 3 {
		fmt.Println("ball!")
		fmt.Println("fourball!")
	}
}
