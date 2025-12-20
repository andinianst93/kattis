// **Ladder**
/// Input: Read wall height h and maximum angle v
/// Convert angle: Degrees to radians
/// Radians = Degrees × (π / 180)
/// Apply trigonometry:
/// sin(v) = h / L
/// L = h / sin(v)
/// Output: Print the ladder length

package ladder

import (
	"fmt"
	"math"
)

func Ladder(){
	var h, v float64

	// read input
	fmt.Scanln(&h, &v)

	// Convert degree to radiants
	angleRadians := v * (math.Pi / 180.0)

	// Calculate ladder length: L = h / sin(v)
	ladderLength := h / math.Sin(angleRadians)

	// Rounded up
	rounded := math.Ceil(ladderLength)

	// Print result
	fmt.Printf("%.0f\n", rounded)
}