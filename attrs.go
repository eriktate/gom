package gom

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
