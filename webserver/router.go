package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ZooService interface {
	Search(text string) []string
}

var r *chi.Mux

func NewWebServer(zoo ZooService) *chi.Mux {

	r = chi.NewRouter()

	fmt.Printf("some number: %s", "abc")
	t := NewTemplate()

	data := map[string]interface{}{
		"HeaderTitle": "Zoo Application",
		"Nav": map[string]interface{}{
			"Pages": []PageItem{
				{Icon: "", Label: "Search Htmx", Link: "/", Page: "/"},
				{Icon: "", Label: "Search Component", Link: "/search-comp", Page: "/search-comp"},
				{Icon: "", Label: "About", Link: "/about", Page: "/about"},
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

	r.Get("/search-comp", func(w http.ResponseWriter, r *http.Request) {
		data["active"] = "/search-comp"
		data["ListTitle"] = "Animals List Component"
		// Optional initialization:
		// data["ListAnimals"] = []string{"Dog", "Rabbit", "Wolf", "Bear"}
		name := "search-comp.page.html"
		err := t.RenderWrapper(w, r, name, data)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		data["active"] = "/about"
		name := "about.page.html"
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

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		searchString := r.URL.Query().Get("searchInput")
		if searchString == "" {
			return
		}
		result := zoo.Search(searchString)
		if len(result) == 0 {
			return
		}
		err := t.RenderComponent(w, "animal-list.html", result)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	r.Get("/search/json", func(w http.ResponseWriter, r *http.Request) {
		searchString := r.URL.Query().Get("searchInput")
		if searchString == "" {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		result := zoo.Search(searchString)
		if len(result) == 0 {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)
	return r
}

func Run(port string) {
	http.ListenAndServe(":"+port, r)
}
