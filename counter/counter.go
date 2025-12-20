package counter

import "fmt"

func Counter(){
	var n int
	fmt.Scanln(&n)

	result := n - 1
	fmt.Println(result)
}