package content

import (
	"testing"
)

func TestCreateFormFromChannel(t *testing.T) {
	type args struct {
		channel <-chan string
	}

	testChannel := make(chan string, 1)
	go send(testChannel)

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ReceivingFromChannel", args: args{channel: testChannel}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()
			if got := CreateFormFromChannel(tt.args.channel); len(got.Fields) != tt.want {
				t.Errorf("CreateFormFromChannel() = %v, want %v", len(got.Fields), tt.want)
			}
		})
	}

	<- testChannel
}

func send(testChannel chan string) {
	defer func() {close(testChannel)}()
	var n = 0
	for n < 5 {
		testChannel <- "test"
		n++
	}
}

func Test_resolveType(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want ValueType
	}{
		{name: "Heading", args: args{line: "##HeadingTitle"}, want: HeadingLine},
		{name: "Input", args: args{line: "[i=checkbox]##Input"}, want: InputLine},
		{name: "Button", args: args{line: "[btn]ButtonWithText"}, want: ButtonLine},
		{name: "Title", args: args{line: "**TitleLine**"}, want: TitleLine},
		{name: "NewLine", args: args{line: ""}, want: NewLine},
		{name: "None", args: args{line: "#*None"}, want: None},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveType(tt.args.line); got != tt.want {
				t.Errorf("resolveType() = %v, want %v", got, tt.want)
			}
		})
	}
}
