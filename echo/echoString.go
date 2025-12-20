package echo

import "fmt"

func EchoString() {
	var word string
	fmt.Scanln(&word)

	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", word)
	}

	fmt.Println(" ")
}
