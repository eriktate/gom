# gom
*gom* stands  for *Go Object Model*. Write HTML in pure Go code, no templating required!

## How does it work?
gom is heavily inspired by [elm-html]( https://github.com/elm-lang/html). After using it
in elm I wanted access to a similarly powerful set of tools serverside when writing Go.

The basic usage pattern looks like this:
```go
import . "github.com/eriktate/gom"

func GimmeSomeHTML() string {
	return Div(
		Attrs(Id("happy-little-div")),
		P(
			Attrs(Class("content", "paragraph")),
			Text("Here's yer HTML!")),
		Input(Attrs(Type("button")))
}
```
