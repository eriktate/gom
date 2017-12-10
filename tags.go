package gom

// Div returns a <div></div>
func Div(attrs Attributes, children ...string) string {
	return wrapTag("div", attrs, children...)
}

// Span returns a <span></span>
func Span(attrs Attributes, children ...string) string {
	return wrapTag("span", attrs, children...)
}

// Ul returns a <ul></ul>
func Ul(attrs Attributes, children ...string) string {
	return wrapTag("ul", attrs, children...)
}

// Li returns a <li></li>
func Li(attrs Attributes, children ...string) string {
	return wrapTag("li", attrs, children...)
}

// Input returns a <input />
func Input(attrs Attributes) string {
	return wrapSelfClosingTag("input", attrs)
}
