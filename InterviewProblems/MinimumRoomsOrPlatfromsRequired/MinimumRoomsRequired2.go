package main

import (
	"go4.org/sort"
	"golang.org/x/exp/errors/fmt"
	"strings"
)

type even struct {
	arrival string
	dept    string
}

func main() {

	plans := []even{
		even{"2E", "3D"},
		even{"3E", "7D"},
		even{"4E", "5D"},
		even{"6E", "8D"},
	}

	planStr := make([]string, 0, len(plans)*2)

	for _, event := range plans {

		planStr = append(planStr, event.arrival)
		planStr = append(planStr, event.dept)

	}
	sort.Strings(planStr)

	roomsInUse, minimumRooms := 0, 0

	for _, val := range planStr {

		if strings.Contains(val, "E") {
			roomsInUse++

			if minimumRooms < roomsInUse {
				minimumRooms = roomsInUse
			}
		} else if strings.Contains(val, "D") {
			roomsInUse--
		}

	}

	fmt.Println(minimumRooms)

}
