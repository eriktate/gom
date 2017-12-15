package gom

import "strings"

// This file contains Attribute definitions as defined in the HTMl
// spec.

// Type returns a 'type' attribute to be applied to
// a tag.
func Type(val string) Attribute {
	return Attr("type", val)
}

// Display returns a 'display' attribute to be applied to a tag.
func Display(val bool) Attribute {
	return Attr("display", val)
}

// Charset returns a 'charset' attribute to be applied to a tag.
func Charset(val string) Attribute {
	return Attr("charset", val)
}

// Lang returns a 'lang' attribute to be applied to a tag.
func Lang(val string) Attribute {
	return Attr("lang", val)
}

// Class returns a 'class' attribute to be applied to a tag.
func Class(vals ...string) Attribute {
	return Attr("class", strings.Join(vals, " "))
}

// Download
func Download(val string) Attribute {
	return Attr("download", val)
}

// Href
func Href(val string) Attribute {
	return Attr("href", val)
}

// HrefLang
func HrefLang(val string) Attribute {
	return Attr("hreflang", val)
}

// Ping
func Ping(vals ...string) Attribute {
	return Attr("ping", strings.Join(vals, " "))
}

// Rel
func Rel(vals ...string) Attribute {
	return Attr("rel", strings.Join(vals, " "))
}

// TargetBlank
func TargetBlank() Attribute {
	return Attr("target", "_blank")
}

// TargetParent deprecated in HTML5.
func TargetParent() Attribute {
	return Attr("target", "_parent")
}

// TargetSelf deprecated in HTML5.
func TargetSelf() Attribute {
	return Attr("target", "_self")
}

// TargetTop deprecated in HTML5.
func TargetTop() Attribute {
	return Attr("target", "_top")
}
