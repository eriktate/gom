# gom
 Write HTML in pure Go code, no templating required!

### !WARNING! gom is currently under heavy development, is missing many tags and attributes, and may change over time. Use at your own risk.

## What is it?
**gom** stands for **Go Object Model**, and it is heavily inspired by [elm-html]( https://github.com/elm-lang/html).
After using it in elm I wanted access to a similarly powerful set of tools serverside when writing Go.

The basic usage pattern looks like this:
```go
import . "github.com/eriktate/gom"

// GimmeDatHTML creates a page that tells you what date/time you're viewing it.
func GimmeDatHTML(dateFormat string) string {
	return Doctype(
		Html(Attrs(Lang("en")),
			Head(nil,
				Meta(Attrs(Charset("utf-8"))),
				Title(nil, "This is a test page!")),
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
```

The pattern of building pages is still very hierarchical and is vaguely reminiscient of something like Pug or HAML.

## Why do I care?
If you've ever written an application in Go that needed to serve up templated HTML (aka the buzzword **server-side-rendering**),
then you know the process for doing so isn't particularly straightforward or robust. You define an HTML file with designated insertion
points that your Go app will ultimately fill with some data. A traditional example of the above looks something like this:
```go
var htmlString = `
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

func TemplateHTML() (string, error) {
	pageData := Page{
		Lang:    "en",
		Charset: "utf-8",
		Time:    time.Now().Format("15:04:05 Monday, January 2, 2006"),
	}

	tmpl, err := template.New("my_page").Parse(htmlString)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(buf, pageData); err != nil {
		return "", err
	}

	page, err := ioutil.ReadAll(buf)
	return string(page), err
}
```
Not only there there more lines, it's definitely not as intuitive what's going on. You need
to be pretty familiar with the `html/template` package and how it works. There's also a strict
separation of layout (`htlmString`) and data (`pageData`) that adds friction to building
dynamic pages.
