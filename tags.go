package gom

// Html returns a <html></html>
func Html5(attrs Attributes, children ...string) string {
	// TODO: This feels dirty, maybe clean it up?
	return "<!DOCTYPE html>" + wrapTag("html", attrs, children...)
}

// Head returns a <head></head>
func Head(attrs Attributes, children ...string) string {
	return wrapTag("head", attrs, children...)
}

// Title returns a <title></title>
func Title(attrs Attributes, children ...string) string {
	return wrapTag("title", attrs, children...)
}

// Body returns a <body></body>
func Body(attrs Attributes, children ...string) string {
	return wrapTag("body", attrs, children...)
}

// Div returns a <div></div>
func Div(attrs Attributes, children ...string) string {
	return wrapTag("div", attrs, children...)
}

// Span returns a <span></span>
func Span(attrs Attributes, children ...string) string {
	return wrapTag("span", attrs, children...)
}

// P returns a <p></p>
func P(attrs Attributes, children ...string) string {
	return wrapTag("p", attrs, children...)
}

// Ul returns a <ul></ul>
func Ul(attrs Attributes, children ...string) string {
	return wrapTag("ul", attrs, children...)
}

// Ol returns a <ol></ol>
func Ol(attrs Attributes, children ...string) string {
	return wrapTag("ol", attrs, children...)
}

// H1 returns a <h1></h1>
func H1(attrs Attributes, children ...string) string {
	return wrapTag("h1", attrs, children...)
}

// H2 returns a <h2></h2>
func H2(attrs Attributes, children ...string) string {
	return wrapTag("h2", attrs, children...)
}

// H3 returns a <h3></h3>
func H3(attrs Attributes, children ...string) string {
	return wrapTag("h3", attrs, children...)
}

// H4 returns a <h4></h4>
func H4(attrs Attributes, children ...string) string {
	return wrapTag("h4", attrs, children...)
}

// H5 returns a <h5></h5>
func H5(attrs Attributes, children ...string) string {
	return wrapTag("h5", attrs, children...)
}

// H6 returns a <h6></h6>
func H6(attrs Attributes, children ...string) string {
	return wrapTag("h6", attrs, children...)
}

// A returns a <a></a>
func A(attrs Attributes, children ...string) string {
	return wrapTag("a", attrs, children...)
}

// Li returns a <li></li>
func Li(attrs Attributes, children ...string) string {
	return wrapTag("li", attrs, children...)
}

// Self closing tags

// Meta returns a <meta />
func Meta(attrs Attributes) string {
	return wrapSelfClosingTag("meta", attrs)
}

// Input returns a <input />
func Input(attrs Attributes) string {
	return wrapSelfClosingTag("input", attrs)
}

// Br returns a <br />
func Br(attrs Attributes) string {
	return wrapSelfClosingTag("br", attrs)
}
