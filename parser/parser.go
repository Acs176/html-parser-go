package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	url  string
	text string
}

var Iterations = 0
var linkList = make([]*Link, 0)

func (l *Link) String() string {
	return "{\n" + l.url + "\n" + l.text + "\n}"
}

func NewLink(url string, text string) *Link {
	return &Link{url, text}
}

func parseAnchor(n *html.Node) string {
	Iterations++
	var toReturn = ""
	if n.Type == html.TextNode {
		toReturn += strings.ReplaceAll(n.Data, "\n", "")

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		toReturn += parseAnchor(c)
	}
	return toReturn
}

func ParseHtml(n *html.Node) []*Link {
	Iterations++

	if n.Type == html.ElementNode && n.Data == "a" {
		var link = NewLink(n.Attr[0].Val, parseAnchor(n))
		linkList = append(linkList, link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseHtml(c)
	}

	return linkList
}
