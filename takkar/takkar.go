package takkar

import "fmt"

func Takkar(){
	// Define variable named Trump and Kim
	var trump, kim int

	// Take user input 
	fmt.Scanln(&trump)
	fmt.Scanln(&kim)

	// if trump is bigger than kim, then MAGA!
	// else if kim is bigger than trum, the FAKE NEWS!
	// else WORLD WAR 3!
	if trump > kim {
		fmt.Println("MAGA!")
	} else if trump < kim {
		fmt.Println("FAKE NEWS!")
	} else {
		fmt.Println("WORLD WAR 3!")
	}
}