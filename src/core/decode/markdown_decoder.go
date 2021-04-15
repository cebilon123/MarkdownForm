package decode

import (
	"core"
	"core/Helper"
	"core/element"
)

type Decoder struct {
	DecodingHandler decoder
}

type MarkdownDecoder struct {
}

func (d *Decoder) Decode(linesChannel chan string) *core.Form {
	return d.DecodingHandler.Decode(linesChannel)
}

func (m *MarkdownDecoder) Decode(linesChannel chan string) *core.Form {
	form := core.NewForm()
	println("Starting decoding")
	for msg := range linesChannel {
		valType := element.ResolveTypeBasedOnTwoFirstCharacters(string(msg[0]) + string(msg[1]))

		if element.IsDoubleSidedType(valType) {
			field := element.Field{Value: element.Value{
				Type: valType,
				Val:  Helper.GetValuesBetweenIndexes(1, len(msg)-2, msg),
			}}

			form.Fields = append(form.Fields, field)
		} else {
			field := element.Field{Value: element.Value{
				Type: valType,
				Val: Helper.GetValuesBetweenIndexes(1, len(msg),msg),
			}}
			form.Fields = append(form.Fields, field)
		}
	}

	println("End of decoding")
	return form
}
