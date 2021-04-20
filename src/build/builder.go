package build

import (
	. "content"
	"github.com/cebilon123/tagger"
)

//Builder expose methods to build form into desired type. For example, into HTML.
type Builder interface {
	Build() string
}

type Output struct {
	Form *ExtendedForm
}

func (o *Output) Build() string {
	form := tagger.Tag{
		Type:  tagger.Form,
		Value: "",
	}

	if form.Children == nil {
		form.Children = make([]tagger.Tag, 0)
	}

	for _, field := range o.Form.Fields {
		form.Children = append(form.Children, buildTag(&field))
	}

	return form.Render()
}

func buildTag(field *ExtendedLineField) tagger.Tag {
	return tagger.Tag{
		Type:     field.Type.ResolveHtmlTag(),
		Value:    field.Value,
		Children: nil, //We are building these tags from 1-deep level list, so there is no children
		Params:   mapParams(field.Params),
	}
}

func mapParams(params []Param) []tagger.Param {
	if params == nil || len(params) < 1 {
		return nil
	}

	mappedParams := make([]tagger.Param, 0)

	for _, param := range params {
		mappedParams = append(mappedParams, tagger.Param{
			Key:   param.Key,
			Value: param.Value,
		})
	}

	return mappedParams
}




