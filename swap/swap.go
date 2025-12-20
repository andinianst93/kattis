package swap

import "fmt"

func Swap() {
	var num int
	fmt.Scan(&num)

	first := num / 10
	second := num % 10

	if first == second {
		fmt.Println(num)
	} else {
		fmt.Printf("%d%d\n", second, first)
	}
}
