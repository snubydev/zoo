package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

var r *chi.Mux

func NewWebServer() *chi.Mux {

	r = chi.NewRouter()

	fmt.Printf("some number: %s", "abc")
	t := NewTemplate()

	data := map[string]interface{}{
		"HeaderTitle": "Zoo Application",
		"Nav": map[string]interface{}{
			"Pages": []PageItem{
				{Icon: "", Label: "Home", Link: "/", Page: "/"},
				{Icon: "", Label: "Graph", Link: "/graph", Page: "/graph"},
			},
		},
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data["active"] = "/"
		name := "home.page.html"
		err := t.RenderWrapper(w, r, name, data)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	r.Get("/graph", func(w http.ResponseWriter, r *http.Request) {
		data["active"] = "/graph"
		name := "graph.page.html"
		r.Header.Set("HX-Trigger-After-Swap", "content-swap")
		err := t.RenderWrapper(w, r, name, data)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	r.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		name := "nav.html"
		err := t.RenderWrapper(w, r, name, data)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)
	return r
}

func Run(port string) {
	http.ListenAndServe(":"+port, r)
}
