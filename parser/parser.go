package parser

type Link struct {
	url  string
	text string
}

func (l *Link) String() string {
	return "{\n" + l.url + "\n" + l.text + "\n}"
}

func NewLink(url string, text string) *Link {
	return &Link{url, text}
}
