package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type View struct {
	Name         string
	FileEnding   string
	BasePath     string
	BaseTemplate string
	Context      interface{}
}

func New(req *http.Request) *View {
	v := &View{
		FileEnding:   ".html",
		BaseTemplate: "base",
		BasePath:     "./template",
	}

	return v
}

func (v *View) Render(w http.ResponseWriter) {
	cwd, _ := os.Getwd()
	tmpl := filepath.Join(cwd, v.BasePath, v.Name+v.FileEnding)
	base := filepath.Join(cwd, v.BasePath, v.BaseTemplate+v.FileEnding)
	t, err := template.ParseFiles(tmpl, base)
	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", v.Context)
}
