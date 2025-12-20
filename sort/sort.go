package sort

import "fmt"

func Sort() {
	var n, m int
	fmt.Scan(&n, &m)

	if n < m {
		fmt.Printf("%d %d\n", n, m)
	} else {
		fmt.Printf("%d %d\n", m, n)
	}
}
