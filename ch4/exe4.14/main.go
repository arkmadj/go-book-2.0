package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
<td><a href='{{.Cache}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplates = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type IssueCache struct {
	Issues         []Issue
	IssuesByNumber map[int]Issue
}

func NewIssueCache(owner, repo string) (ic IssueCache, err error) {
	issues, err := GetIssues(owner, repo)
	if err != nil {
		return
	}
	ic.Issues = issues
	ic.IssuesByNumber = make(map[int]Issue, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
	}
	return
}

func logNonNil(v interface{}) {
	if v != nil {
		log.Print(v)
	}
}

func (ic IssueCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	if len(pathParts) < 3 || pathParts[2] == "" {
		logNonNil(issueListTemplate.Execute(w, ic))
		return
	}
	numStr := pathParts[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number: %s", numStr)))
		if err != nil {
			log.Pr
		}
	}
}
