package main

import (
	"go4.org/sort"
	"golang.org/x/exp/errors/fmt"
)

func main() {

	arrival := []int{900, 940, 1100, 1500, 1800}
	dept := []int{910, 1200, 1120, 1130, 1900, 2000}

	sort.Ints(arrival)
	sort.Ints(dept)

	num := len(arrival)
	i, j, rooms, minimumRooms := 0, 0, 0, 0

	for i < num && j < num {

		if arrival[i] < dept[j] {
			rooms++
			i++
		} else {
			rooms--

			if rooms == 0 {
				rooms = 1
			}

			j++
		}

		if minimumRooms < rooms {
			minimumRooms = rooms
		}

	}

	fmt.Println(minimumRooms)

}
