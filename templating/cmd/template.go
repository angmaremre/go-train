package main

import (
	"os"
	"text/template"
)

var loadedTemplates *template.Template

func init() {
	loadedTemplates = template.Must(template.ParseGlob("*.tmpl"))
}

func main() {
	tmpl, _ := template.ParseFiles("emre.tmpl")
	tmpl.Execute(os.Stdout, os.Args[1])

	loadedTemplates.Execute(os.Stdout, os.Args[1])
}
