package echo

import "fmt"

func Echo(){
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var line string
		fmt.Scanln(&line)
		if i % 2 == 1 {
			fmt.Println(line)
		}
	}
}