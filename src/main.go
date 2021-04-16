package main

import (
	"content"
	"io/ioutil"
)

func main() {
	read := content.NewRead()
	testDat, err:= ioutil.ReadFile("src/test.md")
	if err != nil {
		panic(err)
	}
	go read.ReadContent(string(testDat))
	for line := range read.Channel {
		println(line)
	}
}
