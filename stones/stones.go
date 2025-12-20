package stones

import "fmt"

func Stones() {
	var n int
	fmt.Scanln(&n)

	if n%2 == 1 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
