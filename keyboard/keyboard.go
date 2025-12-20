package keyboard

import (
	"bufio"
	"fmt"
	"os"
)

func Keyboard() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	if len(line) > 0 && line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}

	result := []byte{}

	for i := 0; i < len(line); i++ {
		c := line[i]
		if len(result) == 0 || c != result[len(result)-1] {
			result = append(result, c)
		}
	}
	fmt.Println(string(result))
}
