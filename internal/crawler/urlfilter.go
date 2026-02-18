package crawler

import (
	"net/url"
	"strings"
)

var knownCDNs = []string{
	"cdnjs.cloudflare.com",
	"cdn.jsdelivr.net",
	"unpkg.com",
	"fonts.googleapis.com",
	"fonts.gstatic.com",
}

func ShouldVisit(baseURL, link string, followExternal bool) bool {
	u, err := url.Parse(link)

	if err != nil || u.Scheme == "" {
		return false
	}

	if !IsExternal(baseURL, link) {
		return true
	}

	// external link
	if isCDN(u.Host) {
		return true
	}

	return followExternal
}

func isCDN(host string) bool {
	for _, cdn := range knownCDNs {
		if strings.Contains(host, cdn) {
			return true
		}
	}
	return false
}

func IsExternal(base, link string) bool {
	baseURL, err1 := url.Parse(base)
	linkURL, err2 := url.Parse(link)

	if err1 != nil || err2 != nil {
		return false
	}

	return baseURL.Host != linkURL.Host
}
