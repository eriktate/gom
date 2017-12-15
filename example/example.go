package main

import (
	"log"
	"net/http"
	"time"
	"html/template"

	. "github.com/eriktate/gom"
)

// GimmeDatHTML creates a page that tells you what date/time you're viewing it.
func GimmeDatHTML(dateFormat string) string {
	lang := "en"
	charset := "utf-8"
	title := "This is a test page!"
	return Doctype(
		Html(Attrs(Lang(lang)),
			Head(nil,
				Meta(Attrs(Charset(charset))),
				Title(nil, title)),
			Body(nil,
				Div(Attrs(Class("container", "content")),
					Text("You are viewing this page on "),
					displayDate(dateFormat)))))
}

// example of a 'view' function that can inject data into HTML.
func displayDate(format string) string {
	now := time.Now()
	return Span(Attrs(Class("date")), Text(now.Format(format)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(serveHTML)))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	dateFormat := "15:04:05 Monday, January 2, 2006"
	w.Write([]byte(GimmeDatHTML(dateFormat)))
}

htmlString := `
<!DOCTYPE html>

<html lang="{{.Lang}}">
	<head>
		<meta charset="{{.Charset}}">
		<title>"This is a test page!"</title>
	</head>
	<body>
		<div class="container content">
			You are viewing this page on <span class="date">{.Time}</span>
		</div>
	</body>
</html>
`

type Page struct {
	Lang    string
	Charset string
	Time    string
}

func TemplateHTML() {
	page := Page{
		Lang: "en",
		Charset: "utf-8",
		Time: time.Now().Format("15:04:05 Monday, January 2, 2006"),
	}

}
