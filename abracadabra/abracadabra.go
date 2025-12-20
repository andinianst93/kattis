package abracadabra

import "fmt"

func Abracadabra(){
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d Abracadabra\n", i+1)
	}
}
