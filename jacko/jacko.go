package jacko

import "fmt"

func Jacko(){
	var n, t, m int
	fmt.Scan(&n, &t, &m)

	result := n * t * m
	fmt.Println(result)
}