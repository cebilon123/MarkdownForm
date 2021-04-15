package decode

import (
	"core"
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
	println("Starting decoding")
	for msg := range linesChannel {
		println(msg)
	}

	println("End of decoding")
	return &core.Form{}
}
