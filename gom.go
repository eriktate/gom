package gom

import (
	"fmt"
	"html"
	"log"
	"strconv"
	"strings"
)

// Attribute is a Key/Val pair representing an HTML DOM node attribute.
type Attribute struct {
	Key string
	Val interface{}
}

// Attributes gives us a type alias that we can attack helper functions to.
type Attributes []Attribute

// Attr is a convenience constructor for building Attributes.
func Attr(key string, val interface{}) Attribute {
	return Attribute{Key: key, Val: val}
}

//AttrsString is a convenience constructor for building slices of Attributes by repeating
// Key/Val pairs.
// TODO: Make functions for every attribute type?
func AttrsString(vals ...interface{}) Attributes {
	// Make attrs array big enough to hold the values given
	attrs := make([]Attribute, len(vals)/2)
	idx := 0

	for i := 1; i < len(vals); i += 2 {
		key, ok := vals[i-1].(string)
		if !ok {
			// TODO: Need to figure out how to best deal with this potential failure.
			// Maybe run toString() on the key instead of type asserting?
			log.Println("Key wasn't a string")
		}

		attrs[idx] = Attr(key, vals[i])
		idx++
	}

	if len(vals)%2 != 0 {
		key, ok := vals[len(vals)-1].(string)
		if !ok {
			log.Println("Last key wasn't a string")
		}
		attrs[idx] = Attr(key, "")
	}

	return attrs
}

// Attrs returns a slice of Attributes given a variable number of
// Attribute structs.
func Attrs(vals ...Attribute) Attributes {
	return vals
}

// Text allows you to input renderable text within a DOM node.
// TODO: Should probably make this accept an interface{} and deal with various
// types.
func Text(val string) string {
	return html.EscapeString(val)
}

func (a Attribute) String() string {
	return fmt.Sprintf("%s=%s", html.EscapeString(a.Key), toString(a.Val))
}

func (a Attributes) String() string {
	attrs := make([]string, len(a))
	idx := 0

	for _, attr := range a {
		attrs[idx] = attr.String()
		idx++
	}

	return strings.Join(attrs, " ")
}

func wrapTag(tag string, attrs Attributes, children ...string) string {
	if attrs == nil || len(attrs) == 0 {
		return fmt.Sprintf("<%s>%s</%s>", tag, strings.Join(children, ""), tag)
	}

	return fmt.Sprintf("<%s %s>%s</%s>", tag, attrs.String(), strings.Join(children, ""), tag)
}

func wrapSelfClosingTag(tag string, attrs Attributes) string {
	if attrs == nil || len(attrs) == 0 {
		return fmt.Sprintf("<%s />", tag)
	}

	return fmt.Sprintf("<%s %s />", tag, attrs.String())
}

func toString(val interface{}) string {
	// if str, ok := val.(string); ok {
	// 	return strconv.Quote(str)
	// }

	// TODO: For bools, we should be able to use the shortened form.
	// E.g. display vs display=true
	if b, ok := val.(bool); ok {
		if b {
			return "true"
		}

		return "false"
	}

	// TODO: This seems like it will not be nearly as performant as it should be.
	// Revisit later.
	return strconv.Quote(html.EscapeString(fmt.Sprintf("%v", val)))
}
