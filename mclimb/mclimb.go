package mclimb

import "fmt"

func Mclimb() {
	var n, c, y int

	fmt.Scanln(&n)
	fmt.Scanln(&c)
	fmt.Scanln(&y)

	totalCost := n * c

	fmt.Println(totalCost)
}
