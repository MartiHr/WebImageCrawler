package crawler

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(baseURL, htmlStr string) []string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	var links []string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := base.Parse(attr.Val)
					if err == nil {
						links = append(links, link.String())
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	return links
}
