package triangle

import "fmt"

const rightAngle = 90

func Triangle() {
	var angle1, angle2, angle3 int
	fmt.Scanln(&angle1)
	fmt.Scanln(&angle2)
	fmt.Scanln(&angle3)

	if angle1 == rightAngle || angle2 == rightAngle || angle3 == rightAngle {
		fmt.Println("Ratvinklig Triangel")
	} else if angle1 > rightAngle || angle2 > rightAngle || angle3 > rightAngle {
		fmt.Println("Trubbig Triangel")
	} else {
		fmt.Println("Spetsig Triangel")
	}
}
