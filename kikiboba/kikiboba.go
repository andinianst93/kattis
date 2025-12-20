package kikiboba

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Kikiboba() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := strings.TrimSpace(scanner.Text())

	bCount := 0
	kCount := 0

	for _, char := range word {
		if char == 'b' {
			bCount++
		} else if char == 'k' {
			kCount++
		}
	}

	var category string
	if bCount == 0 && kCount == 0 {
		category = "none"
	} else if bCount > kCount {
		category = "boba"
	} else if kCount > bCount {
		category = "kiki"
	} else {
		category = "boki"
	}

	fmt.Println(category)
}
