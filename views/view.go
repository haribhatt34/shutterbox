package views

import "html/template"

// NewView groups templates together
// it takes variadic parameters
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml",
	)

	//again, below slices withing files are unpacked
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
