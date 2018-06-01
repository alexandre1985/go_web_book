package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}

const tmpl = `Notes are:
{{range .}}
	Title: {{.Title}}, Description: {{.Description}}
{{end}}`

func main() {
	notes := []Note{
		{"text/template", "Templates generate textual output"},
		{"html/template", "Templates generate HTML output"},
	}

	t := template.New("note")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	if err := t.Execute(os.Stdout, notes); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
