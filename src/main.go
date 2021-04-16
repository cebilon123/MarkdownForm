package main

import (
	"content"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	defer elapsed("Decoding markdown into form")()
	read := content.NewRead()
	testDat, err:= ioutil.ReadFile("src/test.md")
	if err != nil {
		panic(err)
	}
	go read.ReadContent(string(testDat))
	form := content.CreateFormFromChannel(read.Channel)
	print(form)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

