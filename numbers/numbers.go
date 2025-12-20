package numbers

import "fmt"

func Numbers() {
	var n int
	fmt.Scanln(&n)

	sum := 0

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		sum += number
	}

	fmt.Println(sum)
}
