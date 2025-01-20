package linkparser

import (
	"errors"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func NewLink(n *html.Node) (*Link, error) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return &Link{Href: a.Val, Text: getTextFromNode(n)}, nil
		}
	}
	return nil, errors.New("no href")
}

func ParseHtml(r io.Reader) []Link {
	// data, err := io.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))

	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	// var nodes map[string][]*html.Node
	nodes := make(map[string][]*html.Node)
	traverse(doc, nodes)
	links := make([]Link, len(nodes["a"]))

	for _, n := range nodes["a"] {
		l, err := NewLink(n)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, *l)
	}
	fmt.Printf("%v", links)

	return nil
}

func traverse(n *html.Node, nodes map[string][]*html.Node) {
	if n.Type == html.ElementNode {
		addNodeToMap(n, nodes)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, nodes)
	}
}

func addNodeToMap(n *html.Node, nodes map[string][]*html.Node) {
	if n.Type == html.ElementNode {
		eleType := n.Data
		if _, ok := nodes[eleType]; ok {
			nodes[eleType] = append(nodes[eleType], n)
		} else {
			nodes[eleType] = []*html.Node{n}
		}
	}
}

func getTextFromNode(n *html.Node) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return c.Data
		}
	}
	return ""
}
