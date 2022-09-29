package main

import (
	"os"
	"text/template"
)

// This is for creating dynamic content or
// showing customized output to the user with the text/template package.
func main() {
	t1 := template.New("t1")
	// . in {{}} is called "actions" to use to dynamically inset content.
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	// template.Must function to panic in case Parse returns an error.
	// This is especially useful for templates initialized in the global scope.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By “executing” the template it generates its text with specific
	// values for its actions.
	// The {{.}} action is replaced by the value passed as a parameter to Execute.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// If the data is a struct, {{.FieldName}} action to access its fields.
	// The fields should be exported to be accessible when a template is executing.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"John Smith"})

	// The same applies to maps; with maps there is no restriction on the case of key names.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else provide conditional execution for templates.
	// A value is considered false if it’s the default value of a type (0, an empty string, nil pointer).
	// This sample demonstrates another feature of templates: using - in actions to trim whitespace.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// With range blocks, we can loop through slices, arrays, maps or channels.
	// Inside the range block {{.}} is set to the current item of the iteration.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
