package round

import (
	"fmt"
	"math"
)

func Round(){
	var n float64
	fmt.Scanln(&n)

	rounded := (int(math.Round(n)))

	fmt.Println(rounded)
}