package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	column "github.com/ahmad/go-book-2.0/ch7/exe7.8"
)

var people = []column.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
	{"Alex", 18},
}

var html = template.Must(template.New("people").Parse(`
<html>
	<body>
		<table>
			<tr>
				<th href="?sort=name">name</th>
				<th><a href="?sort=age">age</a></th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Age}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`))

func main() {
	c := column.NewByColumns(people, 2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "age":
			c.Select(c.LessAge)
		case "name":
			c.Select(c.LessName)
		}
		sort.Sort(c)
		err := html.Execute(w, people)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
}
