/*
 * Given k and total stamped pages s, find the minimum number of trips.
 * To minimize trips, we should maximize the use of k-page stamps (stamp as many pages as possible per trip).
 */

/*
* Use as many k-page trips as possible
* Use 1-page trips for the remainder
* Number of k-page trips = s / k (integer division)
* Number of 1-page trips = s % k (remainder)
* Total trips = (s / k) + (s % k)
 */
package greedy

import "fmt"

func Greedy() {
	var k, s int
	fmt.Scan(&k, &s)

	kPageTrips := s / k
	onePageTrips := s % k

	totalTrips := kPageTrips + onePageTrips

	fmt.Println(totalTrips)
}
