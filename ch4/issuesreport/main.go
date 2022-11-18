package main

import (
	"html/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------
Number: {{.Number}}
Userr: {{.User.Login}}
Title: {{.Tile | printf "%.64s}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
