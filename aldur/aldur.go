package aldur

import "fmt"

func Aldur() {
	var n int
	fmt.Scanln(&n)

	// create an array of size n
	ages := make([]int, n)

	// read n ages into array
	for i := 0; i < n; i++ {
		fmt.Scanln(&ages[i])
	}

	// find minimum
	minAge := ages[0]
	for i := 1; i < n; i++ {
		if ages[i] < minAge {
			minAge = ages[i]
		}
	} 
	// print minimum age
	fmt.Println(minAge)
}
