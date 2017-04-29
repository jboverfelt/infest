package main

import "testing"

func TestParseReopenDateOnly(t *testing.T) {
	const str = "2/2/2012"

	date := parseReopenDate(str)

	if date != str {
		t.Fatalf("Got %s, expected %s", date, str)
	}
}

func TestParseReopenBoth(t *testing.T) {
	const str = "Opened under new owner/name 1/13/2017"
	const expected = "1/13/2017"

	date := parseReopenDate(str)

	if date != expected {
		t.Fatalf("Got %s, expected %s", date, expected)
	}
}

func TestParseReopenJustComment(t *testing.T) {
	const str = "foobar"

	date := parseReopenDate(str)

	if date != "" {
		t.Fatalf("Expected an error")
	}
}
