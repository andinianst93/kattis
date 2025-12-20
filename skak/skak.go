package skak

import "fmt"

func Skak(){
	var r1, c1, r2, c2 int
	fmt.Scanln(&r1, &c1)
	fmt.Scanln(&r2, &c2)

	if r1 == r2 || c1 == c2 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}