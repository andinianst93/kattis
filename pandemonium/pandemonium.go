package pandemonium

import "fmt"

func Pandemonium() {
	var m, t, n int
	fmt.Scanln(&m)
	fmt.Scanln(&t)
	fmt.Scanln(&n)

	// total time = m x n
	totalTime := m * n
	fmt.Println(totalTime)
}
