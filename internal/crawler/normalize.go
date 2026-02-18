package crawler

import (
	"net/url"
	"strings"
)

func NormalizeURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	u.Fragment = "" // remove #fragment

	if u.Path == "" {
		u.Path = "/"
	}

	u.Path = strings.TrimSuffix(u.Path, "/")

	return u.String(), nil
}
