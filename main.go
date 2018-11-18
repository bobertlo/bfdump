package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"gopkg.in/russross/blackfriday.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printNode(n *blackfriday.Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	switch t := n.Type; t {
	case blackfriday.Text:
		text := string(n.Literal)
		text = strings.Replace(text, "\n","\\n",-1)
		fmt.Printf("Text \"%s\"\n",text)
	case blackfriday.Heading:
		fmt.Printf("Heading Level %d\n",n.HeadingData.Level)
	case blackfriday.Link:
		fmt.Printf("Link dst:%s\n",n.LinkData.Destination)
	case blackfriday.Code:
		fmt.Printf("Code: `%s`\n",n.Literal)
	case blackfriday.CodeBlock:
		fmt.Printf("CodeBlock:\n{\n%s}\n",n.Literal)
	default:
		fmt.Println(n.Type);
	}
}

func traverse(n *blackfriday.Node, depth int) {
	printNode(n, depth)
	for c := n.FirstChild; c != nil; c = c.Next {
		traverse(c, depth+2)
	}
}

func main() {
	dat, err := ioutil.ReadFile("test.md")
	check(err)

	m := blackfriday.New(blackfriday.WithExtensions(
		blackfriday.Tables | blackfriday.FencedCode))
	n := m.Parse(dat)

	traverse(n,0);
}
