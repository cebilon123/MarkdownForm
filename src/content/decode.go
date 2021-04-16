package content

import (
	"globals/separators"
	"strings"
)

func CreateFormFromChannel(channel <-chan string) *Form {
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
	if len(line) == 0 {
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
