package content

//Form represents whole form object. Its build from slice of LineField
type Form struct {
	Fields []LineField
}

//NewForm creates new form with fields as slice of size 0
func NewForm() *Form {
	return &Form{
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
