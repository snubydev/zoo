package webserver

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

// https://stackoverflow.com/questions/36617949/how-to-use-base-template-file-for-golang-html-template
//
//go:embed views/*
var tmplFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	funcMap := template.FuncMap{
		// ..
	}
	templates := template.Must(template.New("").
		Funcs(funcMap).
		ParseFS(tmplFS, "views/*.html", "views/partials/*.html"))
	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "views/"+name))
	return tmpl.ExecuteTemplate(w, name, data)
}

func (t *Template) RenderComponent(w io.Writer, name string, data interface{}) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.Parse(fmt.Sprintf(`{{template "%s" .}}`, name)))
	return tmpl.Execute(w, data)
}

func (t *Template) RenderWrapper(w io.Writer, r *http.Request, name string, data interface{}) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		n := strings.Split(name, ".")[0]
		return t.RenderComponent(w, n, data)
	}
	return t.Render(w, name, data)
}
