package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"GO", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
