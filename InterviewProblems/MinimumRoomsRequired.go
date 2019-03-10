package main

import (
	"golang.org/x/exp/errors/fmt"
	"sort"
)

type event struct {
	entry float64
	exit  float64
}

func (slice roomEvent) Len() int {
	return len(slice)
}

func (slice roomEvent) Less(i, j int) bool {
	return slice[i].entry < slice[j].entry
}

func (slice roomEvent) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type roomEvent []event

func main() {

	roomEvents := roomEvent{
		{9.00, 9.10}, {9.40, 12.00}, {9.50, 11.20}, {11.00, 11.30}, {15.00, 19.00}, {18.00, 20.00},
	}
	// more info on sorting : https://yourbasic.org/golang/how-to-sort-in-go/
	sort.Sort(roomEvents)

	roomExitTimings := make([]float64, 0, len(roomEvents))
	var minimumHotelsRequired int

	for i, room := range roomEvents {

		if i == 0 {
			roomExitTimings = append(roomExitTimings, room.exit)
		} else {
			for _, exit := range roomExitTimings {
				if room.entry > exit {
					roomExitTimings = roomExitTimings[1:]
				}
			}

			roomExitTimings = append(roomExitTimings, room.exit)
			sort.Float64s(roomExitTimings)

			if len(roomExitTimings) > minimumHotelsRequired {
				minimumHotelsRequired = len(roomExitTimings)
			}
		}

	}

	fmt.Println(minimumHotelsRequired)

}
