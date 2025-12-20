package spritt

import (
	"bufio"
	"fmt"
	"os"
)

func Spritt() {
	reader := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(reader, &n, &s)

	totalNeeded := 0

	for i := 0; i < n; i++ {
		// bottles needed for classroom i
		var need int
		fmt.Fscan(reader, &need)
		totalNeeded += need
	}

	if totalNeeded <= s {
		fmt.Println("Jebb")
	} else {
		fmt.Println("Neibb")
	}
}
