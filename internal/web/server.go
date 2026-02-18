package web

import (
	"net/http"
)

func Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", SearchPage)
	mux.HandleFunc("/search", SearchImages)

	mux.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))

	return http.ListenAndServe(addr, mux)
}
