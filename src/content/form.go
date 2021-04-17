package content

//BaseForm represents whole form object. Its build from slice of LineField
//its simplified version of normal form object, which is much more advanced
type BaseForm struct {
	Fields []LineField
}

//NewForm creates new form with fields as slice of size 0
func NewForm() *BaseForm {
	return &BaseForm{
		Fields: make([]LineField, 0),
	}
}

type LineField struct {
	Type  ValueType
	Value string
}

type ValueType int

const (
	None ValueType = 0 + iota
	NewLine
	TitleLine
	HeadingLine
	InputLine
	ButtonLine
)

type ExtendedForm struct {
	Fields []ExtendedLineField
}

type ExtendedLineField struct {
	Type ValueType
	Value string
	Params []Param
}

type Param struct {
	Key string
	Value string
}
