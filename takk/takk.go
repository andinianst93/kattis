package takk

import "fmt"

func Takk() {
	var n int
	var name string

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Printf("Takk %s\n", name)
	}
}
