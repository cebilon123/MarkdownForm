package main

import (
	"core/decode"
)

func main() {
	var a = decode.Decoder{
		DecodingHandler: &decode.MarkdownDecoder{
			FileContentText: "## Title",
		},
	}
	a.DecodingHandler.Decode()
}
