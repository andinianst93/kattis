package footballfield

import "fmt"

func FootballField(){
	var n int
	fmt.Scanln(&n)

	result := float64(n) * 0.09144

	fmt.Printf("%.5f\n", result)

}