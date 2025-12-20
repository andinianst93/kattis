package counter

import (
	"bufio"
	"fmt"
	"os"
)

func CountYVowel() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	allVows := 0
	withoutY := 0

	for _, char := range text {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'y' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'Y' {
			allVows++
		}

		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			withoutY++
		}
	}

	fmt.Printf("%d %d\n", withoutY, allVows)
}
