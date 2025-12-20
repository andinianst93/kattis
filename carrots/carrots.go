package carrots

import "fmt"
func Carrots(){
	var numContestants, numCarrots int
	fmt.Scan(&numContestants, &numCarrots)

	for i := 0; i < numContestants; i++ {
		var description string
		fmt.Scanln(&description)
	}

	fmt.Println(numCarrots)
}
