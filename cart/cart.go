package cart

import "fmt"

func Cart() {
	var miles, students, temperature int
	fmt.Scanln(&miles)
	fmt.Scanln(&students)
	fmt.Scanln(&temperature)

	result := miles * 2

	fmt.Printf("%d\n", result)
}
