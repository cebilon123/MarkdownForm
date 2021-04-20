package main

import (
	"build"
	"content"
	"fmt"
	"io/ioutil"
	"time"
)

//TODO After implementing html convert package back to this and use it to convert markdowns into form
func main() {
	defer elapsed("Decoding markdown into form")()

	read := content.NewRead()
	testDat, err := ioutil.ReadFile("src/test.md")

	if err != nil {
		panic(err)
	}

	go read.ReadContent(string(testDat))

	form := content.CreateFormFromChannel(read.Channel).ToExtendedForm()

	output := build.Output{Form: form}
	println(output.Build())
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
