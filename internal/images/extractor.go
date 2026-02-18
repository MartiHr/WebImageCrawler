package images

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractImages(baseURL, htmlStr string) []Image {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	var images []Image

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		// <img> tags
		if n.Type == html.ElementNode && n.Data == "img" {
			var src, alt, title string
			for _, a := range n.Attr {
				switch a.Key {
				case "src":
					src = a.Val
				case "data-src":
					if src == "" {
						src = a.Val
					}
				case "alt":
					alt = a.Val
				case "title":
					title = a.Val
				}
			}

			if src != "" {
				if u, err := base.Parse(src); err == nil {
					images = append(images, Image{
						URL:    u.String(),
						Alt:    alt,
						Title:  title,
						Format: detectFormat(u.String()),
					})
				}
			}
		}

		// Inline SVG
		if n.Type == html.ElementNode && n.Data == "svg" {
			images = append(images, Image{
				Format: "svg",
			})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)
	return images
}

func detectFormat(u string) string {
	u = strings.ToLower(u)
	switch {
	case strings.HasSuffix(u, ".png"):
		return "png"
	case strings.HasSuffix(u, ".jpg"), strings.HasSuffix(u, ".jpeg"):
		return "jpg"
	case strings.HasSuffix(u, ".webp"):
		return "webp"
	case strings.HasSuffix(u, ".svg"):
		return "svg"
	default:
		return "unknown"
	}
}
