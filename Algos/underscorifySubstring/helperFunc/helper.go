package helperFunc

import (
	"strings"
)

func UnderscorifySubstring(str string, substring string) string {

	var locations,combinedLocations []int
	globalIndex := -1

	var underscorifyString []rune

	for i := 0; i < len(str); i++ {

		index := strings.Index(str[i:],substring)
		if index == -1 {
			break
		}

		if globalIndex == -1 {
			globalIndex = index
		} else {
			globalIndex += index + 1
		}

		locations = append(locations, globalIndex, globalIndex+len(substring))
		i = globalIndex
	}

	if len(locations) == 0 {
		return str
	}


	for i := 0; i < len(locations)-1; i++{

		if locations[i] >= locations[i+1] {
			i++
		} else {
			combinedLocations = append(combinedLocations,locations[i])
		}
	}

	if locations[len(locations)-1] != locations[len(locations)-2] {
		combinedLocations = append(combinedLocations,locations[len(locations)-1])
	}

	j := 0
	for i := 0; i < len(str); i++ {
		if j < len(combinedLocations) && i == combinedLocations[j] {
			underscorifyString = append(underscorifyString,'_')
			underscorifyString = append(underscorifyString,rune(str[i]))
			j++
		} else {
			underscorifyString = append(underscorifyString,rune(str[i]))
		}
	}

	if j == len(combinedLocations)-1 {
		underscorifyString = append(underscorifyString,'_')
	}

	return string(underscorifyString)
}


/*func UnderscorifySubstring(str string, substring string) string {
	var locations,combinedLocations []int
	globalIndex := -1

	var underscorifyString []rune

	for i := 0; i < len(str); i++ {

		index := strings.Index(str[i:],substring)
		if index == -1 {
			break
		}

		if globalIndex == -1 {
			globalIndex = index
		} else {
			globalIndex += index + 1
		}

		locations = append(locations, globalIndex, globalIndex+len(substring))
		i = globalIndex
	}


	for i := 0; i < len(locations)-1; i++{

		if locations[i] != locations[i+1] {
			combinedLocations = append(combinedLocations,locations[i])
		} else {
			i++
		}
	}

	if locations[len(locations)-1] != locations[len(locations)-2] {
		combinedLocations = append(combinedLocations,locations[len(locations)-1])
	}

	for i,j := 0,0; i < len(str); i++ {
		if j < len(combinedLocations) && i == combinedLocations[j] {
			underscorifyString = append(underscorifyString,'_')
			underscorifyString = append(underscorifyString,rune(str[i]))
			j++
		} else {
			underscorifyString = append(underscorifyString,rune(str[i]))
		}
	}

	return string(underscorifyString)
}*/