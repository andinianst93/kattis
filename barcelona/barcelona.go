package barcelona

import (
	"bufio"
	"fmt"
	"os"
)

func Barcelona() {
	reader := bufio.NewReader(os.Stdin)

	var n, b int
	fmt.Fscan(reader, &n, &b)

	position := 0

	for i := 1; i <= n; i++ {
		var bag int
		fmt.Fscan(reader, &bag)
		if bag == b && position == 0 {
			position = i
			break
		}
	}

	if position == 1 {
		fmt.Println("fyrst")
	} else if position == 2 {
		fmt.Println("naestfyrst")
	} else {
		fmt.Printf("%d fyrst\n", position)
	}
}
