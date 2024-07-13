package internal

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed map.tmpl
var mapTemplate string

func RenderResults(results []LookupResult) {
	tpl, _ := template.New("map").Parse(mapTemplate)
	_ = tpl.Execute(os.Stdout, results)
}
