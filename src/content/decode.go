package content

func Decode(channel chan<- string) *Form {
	result := NewForm()

	for line := range channel {
		lineType := resolveType(&line)

		lineField := LineField{
			Type:  lineType,
			Value: line,
		}

		result.Fields = append(result.Fields, lineField)
	}

	return result
}

func resolveType(line *string) ValueType {

}
