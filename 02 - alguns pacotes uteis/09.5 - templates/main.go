package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"./public/header.html",
		"./public/content.html",
		"./public/footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Python", 50},
	})

	if err != nil {
		panic(err)
	}

}
