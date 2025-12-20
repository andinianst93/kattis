package blandad

import "fmt"

func Blandad(){
	var n int
	fmt.Scanln(&n)

	hasBeef := false
	hasChicken := false

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		if s == "nautakjot" {
			hasBeef = true
		} else if s == "kjuklingur" {
			hasChicken = true
		}
	}

	if hasBeef && hasChicken {
		fmt.Println("blandad best")
	} else if hasBeef {
		fmt.Println("nautakjot")
	} else {
		fmt.Println("kjuklingur")
	}
}