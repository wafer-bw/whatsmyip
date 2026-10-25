package api

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed page.tmpl
var pageTemplate string

var page = template.Must(template.New("page").Parse(pageTemplate))

// Page renders the homepage with the IP address of the request within the body
// so the page doesn't have to call the API after it loads.
func Page(w http.ResponseWriter, r *http.Request) {
	reply, err := identifyIP(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := page.Execute(w, reply); err != nil {
		log.Printf("rendering page: %v", err)
	}
}
