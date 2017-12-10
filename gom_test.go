package gom_test

import (
	"testing"

	"github.com/eriktate/gom"
)

func Test_Text_Safe(t *testing.T) {
	// SETUP
	testText := "This is some text"

	// RUN
	res := gom.Text(testText)

	// ASSERT
	if res != "This is some text" {
		t.Fatalf("Text altered safe text for some reason")
	}
}

func Test_Text_Unsafe(t *testing.T) {
	// SETUP
	testText := "This is some <unsafe> text"

	// RUN
	res := gom.Text(testText)

	// ASSERT
	if res != "This is some &lt;unsafe&gt; text" {
		t.Fatalf("Text altered safe text for some reason")
	}
}

func Test_Div(t *testing.T) {
	// SETUP
	expectation := "<div id=\"noice\" display=true>Div of Dopeness</div>"

	// RUN
	res := gom.Div(gom.AttrsString("id", "noice", "display", true),
		gom.Text("Div of Dopeness"))

	// ASSERT
	if res != expectation {
		t.Log("Generated DOM does not match expectation")
		t.Logf("Expected: %s", expectation)
		t.Fatalf("Result: %s", res)
	}

}

func Test_Span(t *testing.T) {
	// SETUP
	expectation := "<span id=\"noice\" display=true>A Spectacular Span</span>"

	// RUN
	res := gom.Span(gom.AttrsString("id", "noice", "display", true),
		gom.Text("A Spectacular Span"))

	// ASSERT
	if res != expectation {
		t.Log("Generated DOM does not match expectation")
		t.Logf("Expected: %s", expectation)
		t.Fatalf("Result: %s", res)
	}
}

func Test_Ul_Li(t *testing.T) {
	// SETUP
	expectation := "<ul><li>This</li><li>is</li><li>a test</li></ul>"

	// RUN
	res := gom.Ul(nil,
		gom.Li(nil, gom.Text("This")),
		gom.Li(nil, gom.Text("is")),
		gom.Li(nil, gom.Text("a test")))

	// ASSERT
	if res != expectation {
		t.Log("Generated DOM does not match expectation")
		t.Logf("Expected: %s", expectation)
		t.Fatalf("Result: %s", res)
	}
}

func Test_Input_Type(t *testing.T) {
	// SETUP
	expectation := "<input type=\"text\" />"

	// RUN
	res := gom.Input(gom.Attrs(gom.Type("text")))

	// ASSERT
	if res != expectation {
		t.Log("Generated DOM does not match expectation")
		t.Logf("Expected: %s", expectation)
		t.Fatalf("Result: %s", res)
	}
}
