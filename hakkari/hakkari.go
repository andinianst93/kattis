package hakkari

import (
	"bufio"
	"fmt"
	"os"
)

func Hakkari() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// Store mine position as pairs
	type position struct {
		row, col int
	}
	var mines []position

	for row := 1; row <= n; row++ {
		var line string
		fmt.Fscan(reader, &line)

		for col := 1; col <= m; col++ {
			if line[col-1] == '*' {
				mines = append(mines, position{row, col})
			}
		}
	}

	fmt.Fprintln(writer, len(mines))

	for _, mine := range mines {
		fmt.Fprintln(writer, mine.row, mine.col)
	}
}
