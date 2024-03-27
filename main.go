package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/Acs176/html-parser-go/parser"

	"golang.org/x/net/html"
)

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
	linkList := parser.ParseHtml(doc)
	for _, link := range linkList {
		fmt.Println(link)
	}
	fmt.Println("Iterations: ", parser.Iterations)

}

func getHtmlFileName() string {
	htmlFileName := flag.String("html", "ex5", "name of the html file to be parsed (without extension)")
	flag.Parse()
	if !validFileNames[*htmlFileName] {
		*htmlFileName = "ex1"
	}
	return *htmlFileName
}
