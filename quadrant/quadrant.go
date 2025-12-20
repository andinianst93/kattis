package quadrant

import "fmt"

func Quadrant() {
	var x, y int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	switch {
	case x > 0 && y > 0:
		fmt.Println("1")
	case x < 0 && y > 0:
		fmt.Println("2")
	case x < 0 && y < 0:
		fmt.Println("3")
	case x > 0 && y < 0:
		fmt.Println("4")
	}
}
