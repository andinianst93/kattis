package dna

import (
	"fmt"
	"strings"
)

func Dna(){
	var dna string
	fmt.Scanln(&dna)

	if strings.Contains(dna, "COV") {
		fmt.Println("Veikur!")
	} else {
		fmt.Println("Ekki veikur!")
	}
}