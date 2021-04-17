package content

import (
	"globals/separators"
	"strings"
)

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

func NewExtendedForm() *ExtendedForm {
	return &ExtendedForm{Fields: make([]ExtendedLineField, 0)}
}

type ExtendedLineField struct {
	Type   ValueType
	Value  string
	Params []Param
}

//ToExtendedForm converts BaseForm into ExtendedForm
func (f *BaseForm) ToExtendedForm() *ExtendedForm {
	form := NewExtendedForm()

	for _, field := range f.Fields {
		extendedField := field.toExtendedLineField()
		form.Fields = append(form.Fields, extendedField)
	}

	return form
}

//toExtendedLineField converts LineField into ExtendedLineField. It cleans value from separators, convert found params
//into params array of Param type. If LineField is classic .md line (e.g. ##heading or **title**) it will return
//ExtendedLineField with params as nil. If LineField is an input/btn type it will return ExtendedLineField
//with empty value and slice containing all valid params.
func (lf *LineField) toExtendedLineField() ExtendedLineField {
	if IsParamContainableType(lf.Type) {
		tryMapToParams := TryMapToParams(lf.Value)

		return ExtendedLineField{
			Type:   lf.Type,
			Value:  "",
			Params: tryMapToParams,
		}
	}

	//In this place we are 100% sure that given line is just classic, text .md line and it is not any kind of input/btn
	val := ""
	if len(lf.Value) > 0 {
		val = getCleanedValueBasedOnType(lf.Value, lf.Type)
	}

	return ExtendedLineField{
		Type:   lf.Type,
		Value:  val,
		Params: nil,
	}
}

func getCleanedValueBasedOnType(s string, vt ValueType) string {
	switch valueType := vt; valueType {
	case HeadingLine:
		return strings.TrimRight(s, separators.HeadingSeparatorStart.String())
	case TitleLine:
		return strings.TrimLeft(strings.TrimRight(s, separators.TitleSeparatorStart.String()), separators.TitleSeparatorStart.String())
	default:
		//Just return initial value if there is no type or other problem
		return s
	}
}
