package concatenate

import "fmt"

func Concatenate(){
	var word1, word2 string
	fmt.Scanln(&word1)
	fmt.Scanln(&word2)

	fmt.Printf("%s%s\n", word1, word2)
}