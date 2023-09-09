package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"./public/header.html",
		"./public/content.html",
		"./public/footer.html",
	}

	// t := template.Must(template.New("content.html").ParseFiles(templates...))

	t := template.New("content.html")

	// tipo um bind, quando tu chama ToUpper chama a fn
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Python", 50},
	})

	if err != nil {
		panic(err)
	}

}
