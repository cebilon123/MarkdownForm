package read

type MarkdownReader struct {
	// ReadChannel is channel to which markdown string is read to (returns string lines of markdown)
	ReadChannel chan string
}

func (mr *MarkdownReader) ReadMarkdownContentToChannel(markdownContent string) {
	println("Starting reading")
	defer func() {
		recover()
		close(mr.ReadChannel)
	}()

	var line string
	var wasSingleSidedOperator bool
	var wasDoubleSidedOperator bool
	for i := 0; i < len(markdownContent); i++ {
		if isSingleSidedSeparator(i, markdownContent) {
			line += string(markdownContent[i]) + string(markdownContent[i+1])
			i++
			wasSingleSidedOperator = true
			continue
		}
		if isDoubleSidedSeparator(i, markdownContent) && !wasDoubleSidedOperator {
			line += string(markdownContent[i]) + string(markdownContent[i+1])
			i++
			wasDoubleSidedOperator = true
			continue
		}
		if wasSingleSidedOperator || wasDoubleSidedOperator {
			line += string(markdownContent[i])
		}
		if wasSingleSidedOperator && string(markdownContent[i]) == " " || i == len(markdownContent)-1 {
			mr.ReadChannel <- line
			line = ""
			wasSingleSidedOperator = false
		}
		if wasDoubleSidedOperator && isDoubleSidedSeparator(i, markdownContent) || i == len(markdownContent)-1 {
			line += string(markdownContent[i])
			mr.ReadChannel <- line
			line = ""
			wasDoubleSidedOperator = false
		}
		if wasSingleSidedOperator && i == len(markdownContent)-1 {
			break
		}
	}

	close(mr.ReadChannel)
}

func isSingleSidedSeparator(index int, content string) bool {
	separator := string(content[index]) + string(content[index+1])
	return separator == "##"
}

func isDoubleSidedSeparator(index int, content string) bool {
	separator := string(content[index]) + string(content[index+1])
	return separator == "**"
}
