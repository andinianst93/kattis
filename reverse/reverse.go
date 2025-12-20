package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reverse() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// store numbers in an array
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		numbers[i], _ = strconv.Atoi(scanner.Text())
	}

	// iterate backwards from index n-1 (last index) to 0 (first index)
	for i := n - 1; i >= 0; i-- {
		fmt.Println(numbers[i])
	}
}
