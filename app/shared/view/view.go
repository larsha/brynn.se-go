package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/larsha/fre.la/app/shared/config"
)

type GlobalContext struct {
	Static     string
	Production bool
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
			Static:     config.Static,
			Production: config.Production,
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
