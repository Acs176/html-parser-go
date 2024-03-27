package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Acs176/parser/parser"

	"golang.org/x/net/html"
)

var iterations = 0
var linkList = make([]*parser.Link, 0)
var validFileNames = map[string]bool{
	"ex1": true,
	"ex2": true,
	"ex3": true,
	"ex4": true,
	"ex5": true,
}

func main() {
	htmlFileName := getHtmlFileName()
	file, err := os.Open(htmlFileName + ".html")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	linkList := ParseHtml(doc)
	for _, link := range linkList {
		fmt.Println(link)
	}
	fmt.Println("Iterations: ", iterations)

}

func parseAnchor(n *html.Node) string {
	iterations++
	var toReturn = ""
	if n.Type == html.TextNode {
		toReturn += strings.ReplaceAll(n.Data, "\n", "")

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		toReturn += parseAnchor(c)
	}
	return toReturn
}

func ParseHtml(n *html.Node) []*parser.Link {
	iterations++

	if n.Type == html.ElementNode && n.Data == "a" {
		var link = parser.NewLink(n.Attr[0].Val, parseAnchor(n))
		linkList = append(linkList, link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseHtml(c)
	}

	return linkList
}

func getHtmlFileName() string {
	htmlFileName := flag.String("html", "ex5", "name of the html file to be parsed (without extension)")
	flag.Parse()
	if !validFileNames[*htmlFileName] {
		*htmlFileName = "ex1"
	}
	return *htmlFileName
}
