package counter

import (
	"bufio"
	"fmt"
	"os"
)

func CountCharacter() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	vowelCount := 0

	for _, char := range text {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowelCount++
		}
	}

	fmt.Println(vowelCount)
}
