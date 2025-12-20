package baloon

import "fmt"

func Baloon() {
	var v, a, t int
	fmt.Scan(&v, &a, &t)

	d := float64(v)*float64(t) + 0.5*float64(a)*float64(t)*float64(t)

	fmt.Printf("%.9f\n", d)
}
