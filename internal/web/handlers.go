package web

import (
	"html/template"
	"net/http"
	_ "strings"

	"Homework2/internal/storage"
)

var templates = template.Must(
	template.ParseFiles(
		"internal/web/templates/search.html",
		"internal/web/templates/results.html",
	),
)

var repo *storage.ImageRepository

func Init(r *storage.ImageRepository) {
	repo = r
}

func SearchPage(w http.ResponseWriter, _ *http.Request) {
	templates.ExecuteTemplate(w, "search.html", nil)
}

func SearchImages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	filters := map[string]string{
		"url":        q.Get("url"),
		"filename":   q.Get("filename"),
		"alt":        q.Get("alt"),
		"title":      q.Get("title"),
		"format":     q.Get("format"),
		"min_width":  q.Get("min_width"),
		"max_width":  q.Get("max_width"),
		"min_height": q.Get("min_height"),
		"max_height": q.Get("max_height"),
	}

	results, err := repo.SearchAll(filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "results.html", results)
}
