package triangle

import "fmt"

func Area() {
	var height, base int
	fmt.Scan(&height, &base)

	area := (float64(height) * float64(base)) / 2.0

	fmt.Printf("%.1f\n", area)
}
