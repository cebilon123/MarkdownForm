package separators

type Separator int

const (
	HeadingSeparatorStart Separator = iota
	TitleSeparatorStart
	NewLineSeparatorStart
	InputSeparatorStart
	BtnSeparatorStart
)

func (s Separator) String() string {
	return [...]string{"##", "**", "\n", "[i=", "[btn"}[s]
}

//IsBodyInvolvedSeparator checks if given separator can have body with additional parameters
func (s Separator) IsBodyInvolvedSeparator() bool {
	return s == InputSeparatorStart || s == BtnSeparatorStart
}
