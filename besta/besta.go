package besta

import (
	"fmt"
	"math"
)

func Besta(){
	var n int
	fmt.Scanln(&n)

	max_fun := math.MinInt32
	max_name := ""

	for i := 0; i < n; i++{
		var name string
		var fun int
		fmt.Scanln(&name, &fun)
		if fun > max_fun{
			max_fun = fun
			max_name = name
		}
	}

	fmt.Println(max_name)
}