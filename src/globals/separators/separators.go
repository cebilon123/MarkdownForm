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
