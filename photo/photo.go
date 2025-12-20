package photo

import "fmt"

func Photo() {
	var w, l int
	fmt.Scanln(&w)
	fmt.Scanln(&l)

	result := w * l
	fmt.Println(result)
}
