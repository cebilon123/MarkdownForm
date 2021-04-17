package content

import (
	"strings"
)

type Param struct {
	Key   string
	Value string
}

//IsParamContainableType checks if given value type can contain any params
func IsParamContainableType(valueType ValueType) bool {
	return valueType == InputLine || valueType == ButtonLine
}

//TryMapToParams tries to map given param string to Param slice. If there is empty string or there is no valid param
//it will return nil
func TryMapToParams(s string) []Param {
	params := make([]Param, 0)

	if len(s) == 0 || !strings.Contains(s, "[") || !strings.Contains(s, "]"){
		return nil
	}

	str := strings.TrimLeft(strings.TrimRight(s, "]"), "[")
	for _, param := range strings.Split(str, ",") {
		param = strings.ReplaceAll(param, " ", "")

		paramKeyValue := strings.Split(param, "=")

		if len(paramKeyValue) < 2 {
			continue
		}

		params = append(params, Param{
			Key:   paramKeyValue[0],
			Value: paramKeyValue[1],
		})
	}

	return params
}
