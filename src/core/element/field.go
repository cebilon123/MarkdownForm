package element

// Field represents abstraction of field. It means, that it can be input, image, video etc..
type Field struct {
	Position
	Body
}

func CreateField(position Position, body Body) *Field {
	return &Field{Position: position, Body: body}
}

// Position is made of x and y values as coordinates.
type Position struct {
	X int
	Y int
}

// Body of the field
type Body struct {
	Type  BodyType
	Value Value
}

type BodyType int

const (
	BodyTypeNone          BodyType = iota
	BodyTypeInputText     BodyType = iota
	BodyTypeInputCheckbox BodyType = iota
	BodyTypeImage         BodyType = iota
	BodyTypeVideo         BodyType = iota
)

type Value struct {
	Type ValueType
	Val  string
}

type ValueType int

const (
	ValueTypeNone    ValueType = iota
	ValueTypeHeading ValueType = iota
)
