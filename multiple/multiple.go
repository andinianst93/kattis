package multiple

import "fmt"

func Multiple() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= 12; i++ {
		multiple := n * i
		fmt.Println(multiple)
	}
}
