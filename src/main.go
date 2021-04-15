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

	go mr.ReadMarkdownContentToChannel("##dupka **XD** **lol**")
	decoder.Decode(mr.ReadChannel)
	<- mr.ReadChannel
}
