package decode

import (
	"core"
	"fmt"
	"reflect"
)

type Decoder struct {
	DecodingHandler decoder
}

type MarkdownDecoder struct {
	FileContentText string
}

func (m *MarkdownDecoder) Decode() *core.Form {

	if len(m.FileContentText) == 0 {
		panic(fmt.Sprintf("No body to decode from, type: %s", reflect.TypeOf(m).Elem().Name()))
		return nil
	}
	panic("implement me")
}
