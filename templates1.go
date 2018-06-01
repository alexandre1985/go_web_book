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

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}
`

func main() {
	note := Note{"text/template", "Template generates textual output"}

	t := template.New("note0")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Fatal: ", err)
		return
	}
}
