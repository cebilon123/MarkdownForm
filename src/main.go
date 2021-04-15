package main

import "core/read"

func main() {
	mr := read.MarkdownReader{
		ReadChannel: make(chan string),
	}

	go mr.ReadMarkdownContentToChannel("**test** ##dupka **XD** **lol")
	for elem := range mr.ReadChannel {
		println(elem)
	}
}
