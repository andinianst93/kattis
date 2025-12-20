package addition

import "fmt"

func AdditionTrouble() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	add := a + b

	if add == c {
		fmt.Println("correct!")
	} else {
		fmt.Println("wrong!")
	}
}
