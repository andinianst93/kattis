package gluttonous

import "fmt"

func Gluttonous(){
	var n, m int
	var questionMark string
	fmt.Scan(&n, &questionMark, &m)

	if n > m {
		fmt.Println(">")
	} else if n < m {
		fmt.Println("<")
	} else {
		fmt.Println("Goggi svangur!")
	}
}
