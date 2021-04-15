package element

// Field represents abstraction of field. It means, that it can be input, image, video etc..
type Field struct {
	Value Value
}

func NewField(value Value) *Field {
	return &Field{Value: value}
}

type Value struct {
	Type ValueType
	Val  string
}

type ValueType int

const (
	ValueTypeNone    ValueType = iota
	ValueTypeNewLine ValueType = iota
	ValueTypeTitle   ValueType = iota
	ValueTypeHeading ValueType = iota
)

func ResolveTypeBasedOnTwoFirstCharacters(chars string) ValueType {
	switch chr := chars; chr {
	case "##":
		return ValueTypeTitle
	case "\n":
		return ValueTypeNewLine
	case "**":
		return ValueTypeHeading
	default:
		return ValueTypeNone
	}
}

func IsDoubleSidedType(valType ValueType) bool {
	return valType == ValueTypeHeading
}



