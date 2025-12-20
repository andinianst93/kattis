package blazes

import "fmt"

func Blazes() {
	var n int
	fmt.Scanln(&n)

	totalBlazes := 0

	for i := 0; i < n; i++ {
		var blazesInSection int
		fmt.Scanln(&blazesInSection)
		totalBlazes += blazesInSection
	}

	fmt.Println(totalBlazes)
}
