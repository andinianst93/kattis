package haystack

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HayStack() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	var needle string
	fmt.Fscan(reader, &needle)

	var haystack string
	fmt.Fscan(reader, &haystack)

	if strings.Contains(haystack, needle) {
		fmt.Println("Unnar fann hana!")
	} else {
		fmt.Println("Unnar fann hana ekki!")
	}
}
