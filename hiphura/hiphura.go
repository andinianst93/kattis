package hiphura

import "fmt"


func Hiphura(){
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	for i := 0; i < age; i++ {
		fmt.Printf("Hipp hipp hurra, %s!\n", name)
	}}