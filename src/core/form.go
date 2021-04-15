package core

import . "core/element"

type Form struct {
	Fields []Field
}

func NewForm() *Form {
	return &Form{Fields: make([]Field,0)}
}
