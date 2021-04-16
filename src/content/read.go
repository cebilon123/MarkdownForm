package content

import (
	"bufio"
	"strings"
)

type Reader interface {
	//ReadContent should content content into channel which is being received somewhere
	ReadContent(content string)
}

//Read is responsible for hold of the channel in the midst of which each line being sent.
type Read struct {
	//Channel in the midst of which each of content line being sent
	Channel chan string
}

//NewRead creates new Read object with channel buffer size of 1
func NewRead() *Read {
	channel := make(chan string, 1)
	return &Read{
		Channel: channel,
	}
}

//ReadContent reads given content string line by line and send it to channel
func (r Read) ReadContent(content string) {
	defer func() {
		close(r.Channel)
	}()

	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		r.Channel <- scanner.Text()
	}
}


