package crawler

import "sync"

type URLSet struct {
	mu  sync.Mutex
	set map[string]struct{}
}

func NewURLSet() *URLSet {
	return &URLSet{
		set: make(map[string]struct{}),
	}
}

func (u *URLSet) Add(url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.set[url]; exists {
		return false
	}

	u.set[url] = struct{}{}
	return true
}
