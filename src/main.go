package main

import (
	"core/decode"
	"core/read"
)

func main() {
	mr := read.MarkdownReader{
		ReadChannel: make(chan string, 1),
	}

	decoder := decode.Decoder{DecodingHandler: &decode.MarkdownDecoder{}}
	go mr.ReadMarkdownContentToChannel("##hejka **XD** **lol**")
	form := decoder.Decode(mr.ReadChannel)
	println(form.Fields)
	<- mr.ReadChannel
}
