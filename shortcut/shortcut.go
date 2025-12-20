package shortcut

import "fmt"

func Shortcut(){
	var n int

	fmt.Scanln(&n)

	// result = (N + 5) × 3 - 10
	result := (n + 5) * 3 - 10

	fmt.Println(result)

}