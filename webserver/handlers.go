package webserver

import (
	"encoding/json"
	"net/http"
	"github.com/a-h/templ"
	"zoo/webserver/components"
	"zoo/webserver/models"
)
	
var data = models.Data{
	Active: "/",
	HeaderTitle: "Zoo Application",
	Pages: []models.PageItem{
		{Icon: "", Label: "Search Htmx", Link: "/", Page: "/"},
		{Icon: "", Label: "Search Component", Link: "/search-comp", Page: "/search-comp"},
		{Icon: "", Label: "About", Link: "/about", Page: "/about"},
	},
} 

func Home(w http.ResponseWriter, r *http.Request) {
	data.Active = "/";
	renderHelper(w, r, components.Home(), data);
}

func About(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("HX-Trigger-After-Swap", "content-swap")
	data.Active = "/about";
	renderHelper(w, r, components.About(), data);
}

func SearchComp(w http.ResponseWriter, r *http.Request) {
	data.Active = "/search-comp"
	data.ListTitle = "Animals List Component"
	// Optional initialization:
	data.ListAnimals = []string{"Dog", "Rabbit", "Wolf", "Bear"}
	renderHelper(w, r, components.SearchComp(data), data);
}

func Search(searchFunc func(s string)[]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchString := r.URL.Query().Get("searchInput")
		if searchString == "" {
			return
		}
		result := searchFunc(searchString)
		if len(result) == 0 {
			return
		}
		components.AnimalList(result).Render(r.Context(), w);
	}
}

func SearchJson(searchFunc func(s string)[]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchString := r.URL.Query().Get("searchInput")
		if searchString == "" {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		result := searchFunc(searchString)
		if len(result) == 0 {
			_ = json.NewEncoder(w).Encode([]string{})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	}
}

func renderHelper(w http.ResponseWriter, r *http.Request, comp templ.Component, data models.Data) {
	if (len(r.Header.Get("HX-Request")) > 0) {
		comp.Render(r.Context(), w);
		return
	}
	components.Layout(data, comp).Render(r.Context(), w);
}
