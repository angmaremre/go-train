package main

import (
	"os"
	"text/template"
)

type Marka struct {
	Id        int
	Name      string
	Is_active bool
}

type Markalar []Marka

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("*.emre"))
}

func main() {
	marka1 := Marka{Id: 1, Name: "Adisas", Is_active: true}
	marka2 := Marka{Id: 2, Name: "Nike", Is_active: false}

	data := Markalar{}
	data = append(data, marka1)
	data = append(data, marka2)

	fileStar, _ := os.Create("index.html")
	templates.ExecuteTemplate(fileStar, "index.emre", data)
}
