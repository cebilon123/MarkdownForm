package params

import (
	"content"
	"strings"
)

type Param struct {
	Key string
	Value string
}

//IsParamContainableType checks if given value type can contain any params
func IsParamContainableType(valueType content.ValueType) bool {
	return valueType == content.InputLine || valueType == content.ButtonLine
}

//TryMapToParams tries to map given param string to Param slice. If there is empty string or there is no valid param
//it will return nil
func TryMapToParams(s string) []Param {
	if len(s) == 0 {
		return nil
	}

	str := strings.TrimLeft(strings.TrimRight(s,"["), "]")
}