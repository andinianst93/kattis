package metronome

import "fmt"

func Metronome() {
	var tick int
	fmt.Scanln(&tick)

	tickFloat := float32(tick)
	result := tickFloat / 4.0

	fmt.Printf("%.2f\n", result)
}