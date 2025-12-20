package dagatal

import "fmt"

func Dagatal(){
	var m int
	fmt.Scanln(&m)

	switch m {
	case 2:
		fmt.Println("28")
	case 4, 6, 9, 11:
		fmt.Println("30")
	default:
		fmt.Println("31")
	}
}