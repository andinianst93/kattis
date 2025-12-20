package hiss

import "fmt"

func Hiss() {
	var word string
	fmt.Scanln(&word)

	// Loop through the string until second-to-last character
	for i := 0; i < len(word)-1; i++ {
		if word[i] == 's' && word[i+1] == 's' {
			fmt.Println("hiss")
			return
		}
	}

	fmt.Println("no hiss")
}
