package traffic

import (
	"bufio"
	"fmt"
	"os"
)

func Traffic() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	count := 0
	totalCells := n * m

	for i := 0; i < m; i++ {
		var lane string
		fmt.Fscan(reader, &lane)

		for _, cell := range lane {
			if cell == '.' {
				count++
			}
		}
	}
	proportion := float64(count) / float64(totalCells)
	fmt.Println(proportion)
}
