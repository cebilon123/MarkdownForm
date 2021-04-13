package element

// Field represents abstraction of field. It means, that it can be input, image, video etc..
type Field struct {
	Position
	Body
}

// Position is made of x and y values as coordinates.
type Position struct {
	X int
	Y int
}

// Body of the field
type Body struct {
	Type  BodyType
	Value string
	Label string
}

type BodyType int
const (
	None          BodyType = iota
	InputText     BodyType = iota
	InputCheckbox BodyType = iota
	Image         BodyType = iota
	Video         BodyType = iota
)
