package qaly

import "fmt"

func Qaly() {
	var n int
	fmt.Scanln(&n)

	totalQaly := 0.0

	for i := 0; i < n; i++ {
		var quality, years float64
		fmt.Scan(&quality, &years)
		currentQaly := quality * years
		totalQaly += currentQaly
	}

	fmt.Printf("%.3f\n", totalQaly)
}
