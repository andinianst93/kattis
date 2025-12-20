package fifa

import "fmt"


func Fifa(){
	const year = 2022
	var n, k int

	fmt.Scanln(&n)
	fmt.Scanln(&k)

	currentYear := year + (n / k)
	fmt.Println(currentYear)
}