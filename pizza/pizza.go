package pizza

import "fmt"

func Pizza(){
	// Pizza is n, Number of people is m
	// How many slices will be left over?
	// Leftover slices = n mod m
	// Or written more formally:
	// r = n - m⌊n/m⌋
	var n, m int

	fmt.Scanln(&n)
	fmt.Scanln(&m)

	leftover := n % m

	fmt.Println(leftover)
}