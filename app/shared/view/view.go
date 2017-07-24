package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/larsha/brynn.se-go/app/shared/config"
)

type GlobalContext struct {
	StaticFolder string
	Production   bool
	Cachebust    string
}

type View struct {
	Name          string
	FileEnding    string
	BasePath      string
	BaseTemplate  string
	Context       interface{}
	GlobalContext GlobalContext
	Status        int
}

type ViewContext struct {
	Context interface{}
	Globals GlobalContext
}

func New(req *http.Request) *View {
	config := config.Get()

	v := &View{
		FileEnding:   ".html",
		BaseTemplate: "base",
		BasePath:     "./template",
		GlobalContext: GlobalContext{
			StaticFolder: config.StaticFolder,
			Production:   config.Production,
			Cachebust:    config.Cachebust,
		},
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

	view := ViewContext{
		v.Context,
		v.GlobalContext,
	}

	if v.Status == 0 {
		v.Status = http.StatusOK
	}

	w.WriteHeader(v.Status)
	t.ExecuteTemplate(w, "base", view)
}
