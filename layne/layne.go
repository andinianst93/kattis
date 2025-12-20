package layne

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Layne() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	cleanedEmail := cleanEmail(email)
	fmt.Println(cleanedEmail)
}

func cleanEmail(email string) string {
	var result strings.Builder
	result.Grow(len(email))

	for i := 0; i < len(email); i++ {
		if email[i] != ' ' {
			result.WriteByte(email[i])
		}
	}
	return result.String()
}
