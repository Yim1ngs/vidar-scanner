package basework

import (
	//"fmt"
	"strings"

	"golang.org/x/net/html"
)

var tagsToRemove = map[string]bool{
	"script": true,
	"style":  true,
}

func removeScriptNodes(n *html.Node, tags map[string]bool) {
	if n == nil {
		return
	}

	for child := n.FirstChild; child != nil; {
		next := child.NextSibling

		if child.Type == html.ElementNode && tags[child.Data] {
			n.RemoveChild(child)
		} else {
			removeScriptNodes(child, tags)
		}
		child = next
	}
}

func extractText(n *html.Node, texts *[]string) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		content := strings.TrimSpace(n.Data)
		if content != "" {
			*texts = append(*texts, content)
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		extractText(child, texts)
	}
}

func HTMLPreprocess(HTMLData string) (string, error) {
	doc, err := html.Parse(strings.NewReader(HTMLData))
	if err != nil {
		return "", err
	}

	removeScriptNodes(doc, tagsToRemove)

	var texts []string
	extractText(doc, &texts)

	output := strings.Join(texts, " ")

	//fmt.Println(output)
	return output, nil
}
