package cinema

import "fmt"

func Cinema(){
	var x, y, z int

	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)

	result := calculateLatestStartTime(x, y, z)
	fmt.Println(result)
}

func calculateLatestStartTime(x, y, z int) int {
	// x = minutes from Hannes to Arnar
    // y = minutes from Arnar to cinema
    // z = minute of day when film starts
	totalTravelTime := x + y
	latestStart := z - totalTravelTime

	return latestStart
}