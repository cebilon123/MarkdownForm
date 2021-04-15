package Helper

func GetValuesBetweenIndexes(first int, second int, text string) string{
	var line string
	for index, char := range text {
		if index > first && index < second {
			line += string(char)
		}
	}

	return line
}
