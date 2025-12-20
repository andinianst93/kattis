package draga

import "fmt"

func Draga(){
	// n = window covered by curtains, m = window curtains that have been drawn open
	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	result := n - m
	fmt.Println(result)
}