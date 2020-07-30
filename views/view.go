package views

import "html/template"

// NewView groups templates together
// it takes variadic parameters
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	//again, below slices withing files are unpacked
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View type defined
type View struct {
	Template *template.Template
}
