package bergmal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func Bergmal(){
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", line)

}

