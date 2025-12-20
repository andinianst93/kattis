package flatterLand

import "fmt"

func FlatterLand() {
	var n, d int
	fmt.Scanln(&n)
	fmt.Scanln(&d)

	numberOfGaps := n - 1
	totalDistance := numberOfGaps * d

	fmt.Println(totalDistance)
}
