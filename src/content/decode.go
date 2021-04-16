package content

import (
	"globals/separators"
	strings "strings"
)

func Decode(channel chan<- string) *Form {
	results := NewForm()

	for line := range channel {
		lineType := resolveType(line)

		lineField := LineField{
			Type:  lineType,
			Value: line,
		}

		results.Fields = append(results.Fields, lineField)
	}

	return results
}

func resolveType(line string) ValueType {
	if strings.HasPrefix(line, separators.NewLineSeparatorStart.String()) {
		return NewLine
	}
	if strings.HasPrefix(line, separators.BtnSeparatorStart.String()) {
		return ButtonLine
	}
	if strings.HasPrefix(line, separators.HeadingSeparatorStart.String()) {
		return HeadingLine
	}
	if strings.HasPrefix(line, separators.InputSeparatorStart.String()) {
		return InputLine
	}
	if strings.HasPrefix(line, separators.TitleSeparatorStart.String()) {
		return TitleLine
	}

	return None
}
