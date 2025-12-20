package wind

import (
	"bufio"
	"fmt"
	"os"
)

func Wind() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var w, n int
	fmt.Fscan(reader, &w)
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var name string
		var maxSafe int
		fmt.Fscan(reader, &name, &maxSafe)

		if w <= maxSafe {
			fmt.Fprintf(writer, "%s opin\n", name)
		} else {
			fmt.Fprintf(writer, "%s lokud\n", name)
		}
	}
}
