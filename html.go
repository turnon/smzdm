package main

//go:generate sh templates/generate.sh

import (
	"html/template"
	"io"
	"os"

	"github.com/turnon/smzdm/templates"
)

type html struct {
	resultSet
}

func (out *html) print(ws ...io.Writer) {
	var w io.Writer
	if len(ws) == 0 {
		w = os.Stdout
	} else {
		w = ws[0]
	}

	t := template.New("a")
	t.Parse(templates.Templates["templates/rich"])
	now := out.createdAt.Format("06-01-02 15:04:05")
	t.Execute(w, map[string]interface{}{"data": out.searches, "now": now})
}
