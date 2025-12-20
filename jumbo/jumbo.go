package jumbo

import "fmt"

func Jumbo() {
	var n int
	fmt.Scanln(&n)

	totalSum := 0

	for i := 0; i < n; i++ {
		var rodLength int
		fmt.Scanln(&rodLength)
		totalSum += rodLength
	}

	numberOfFusions := n - 1
	finalLength := totalSum - numberOfFusions
	fmt.Println(finalLength)
}
