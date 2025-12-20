package counter

import (
	"bufio"
	"fmt"
	"os"
)

func CountLetters() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	count := 0

	for _, char := range text {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			count++
		}
	}

	fmt.Println(count)
}
