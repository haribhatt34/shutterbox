package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

// NewView groups templates together
// it takes variadic parameters
func NewView(layout string, files ...string) *View {
	// layoutfiles() will return a slice, but since
	// append take variadic parameters, ... will
	// unpack the slice and convert it strings.
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View type defined
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefiend layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles returns a slice of string reprenting
// the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
