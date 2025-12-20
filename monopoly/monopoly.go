package monopoly

import "fmt"

func Monopoly() {
	waysMap := map[int]int{
		2:  1,
		3:  2,
		4:  3,
		5:  4,
		6:  5,
		7:  6,
		8:  5,
		9:  4,
		10: 3,
		11: 2,
		12: 1,
	}

	var n int
	fmt.Scan(&n)

	totalWays := 0
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		totalWays += waysMap[d]
	}

	probability := float64(totalWays) / 36.0

	fmt.Println(probability)
}
