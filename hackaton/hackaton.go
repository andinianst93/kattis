package hackaton

import "fmt"

func Hackaton() {
	var n, moneySpent, avgPrice int
	fmt.Scanln(&n)
	fmt.Scanln(&moneySpent)
	fmt.Scanln(&avgPrice)

	totalPrices := n * avgPrice

	fmt.Printf("%d\n", totalPrices)
}
