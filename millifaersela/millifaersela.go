package millifaersela

import "fmt"

func Millifaersela(){
	var monnei, fjee, dolladollabilljoll int

	fmt.Scanln(&monnei)
	fmt.Scanln(&fjee)
	fmt.Scanln(&dolladollabilljoll)

	if monnei < fjee && monnei < dolladollabilljoll {
		fmt.Println("Monnei")
	} else if fjee < monnei && fjee < dolladollabilljoll {
		fmt.Println("Fjee")
	} else {
		fmt.Println("Dolladollabilljoll")
	}
}