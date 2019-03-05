package helperFunc

func Shift(str string, key int) string {

	var temp []rune
	key = key % 26

	for _,char := range str {
		newChar := char + int32(key)
		if newChar > 122 {
			temp = append(temp,(newChar%122)+96)
		} else {
			temp = append(temp,newChar)
		}
	}

	return string(temp)

}
