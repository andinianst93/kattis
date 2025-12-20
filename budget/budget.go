package budget

import (
	"bufio"
	"fmt"
	"os"
)

func Budget() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	total := 0

	for i := 0; i < n; i++ {
		var description string
		var amount int

		fmt.Fscan(reader, &description)
		fmt.Fscan(reader, &amount)

		total += amount
	}

	if total > 0 {
		fmt.Println("Usch, vinst")
	} else if total == 0 {
		fmt.Println("Lagom")
	} else {
		fmt.Println("Nekad")
	}
}
