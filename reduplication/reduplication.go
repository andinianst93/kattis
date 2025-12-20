package reduplication

import "fmt"

func Reduplication(){
	var word string
	var number int

	fmt.Scanln(&word)
	fmt.Scanln(&number)

	for i := 0; i < number; i++ {
		fmt.Printf("%s",word)
	}

	fmt.Println()
}