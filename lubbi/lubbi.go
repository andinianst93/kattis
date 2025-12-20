package lubbi

import (
	"fmt"
)

func Lubbi() {
	var word string
	fmt.Scanln(&word)

	if len(word) > 0 {
		firstCharacter := string(word[0])
		fmt.Printf("%s\n", firstCharacter)
	}
}
