package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

var issuelist = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	var issues = newIssues()
	if err := issuelist.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func newIssues() IssuesSearchResult {
	var issues IssuesSearchResult
	issues.TotalCount = 1
	var items []*Issue

	items = append(items, &Issue{
		Number:  1,
		HTMLURL: "html 1",
		Title:   "title 1",
		State:   "state 1",
		User: &User{
			Login:   "login 1",
			HTMLURL: "html 1",
		},
		CreatedAt: time.Now(),
		Body:      "body 1",
	})
	items = append(items, &Issue{
		Number:  1,
		HTMLURL: "html 2",
		Title:   "title 2",
		State:   "state 2",
		User: &User{
			Login:   "login 2",
			HTMLURL: "html 2",
		},
		CreatedAt: time.Now().Add(time.Duration(-48) * time.Hour),
		Body:      "body 2",
	})
	issues.Items = items
	return issues
}
