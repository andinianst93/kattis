package daoi

import "fmt"


func Daoi(){
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	d := b * b - 4 * a * c

	if d > 0 {
		fmt.Println("2")
	} else if d == 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}