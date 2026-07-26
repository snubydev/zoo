package webserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
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

	r.Get("/", Home)

	r.Get("/search-comp", SearchComp)
	
	r.Get("/about", About)
	
	r.Get("/search", Search(zoo.Search)) 
	
	r.Get("/search/json", SearchJson(zoo.Search))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)
	return r
}

func Run(port string) {
	fmt.Printf("Web webserver starting on port: %s \n", port)
	http.ListenAndServe(":"+port, r)
}
