package stringutil

import (
	"testing"
)

func TestKeep(T *testing.T) {

	type input struct {
		source, keep string
	}

	table := map[input]string{
		input{"987-654-3210", Digits}:                            "9876543210",
		input{"(312) 384-7122", Digits}:                          "3123847122",
		input{"This string has the numbers 3, 4, and 5", Digits}: "345",
		input{"No numbers here", Digits}:                         "",
		input{"123 West Main Street", Digits}:                    "123",
		input{"123 West Main Street", Alpha}:                     "WestMainStreet",
	}

	for in, want := range table {
		got := Keep(in.source, in.keep)

		if got != want {
			T.Errorf("Got: %s, Wanted: %s", got, want)
		}
	}
}

func TestStrip(T *testing.T) {

	type input struct {
		source, keep string
	}

	table := map[input]string{
		input{"987-654-3210", Punctuation}:                                   "9876543210",
		input{"(312) 384-7122", Punctuation}:                                 "312 3847122",
		input{"This.... 'string' has the numbers 3, 4, and 5!", Punctuation}: "This string has the numbers 3 4 and 5",
		input{"No numbers here", Punctuation}:                                "No numbers here",
		input{"123 West Main Street", NonAlphaNum}:                           "123",
		input{"123 West Main Street", NonAlphaNum}:                           "123WestMainStreet",
	}

	for in, want := range table {
		got := Strip(in.source, in.keep)

		if got != want {
			T.Errorf("Got: %s, Wanted: %s", got, want)
		}
	}
}

func TestCamelCaseToUnderscored(t *testing.T) {

	testCases := make(map[string]string)

	testCases[""] = ""
	testCases["a"] = "a"
	testCases["A"] = "a"
	testCases["Foo"] = "foo"
	testCases["FooBar"] = "foo_bar"
	testCases["fooBar"] = "foo_bar"
	testCases["thisExample withSpaces"] = "this_example with_spaces"

	for input, expected := range testCases {
		actual := CamelCaseToUnderscored(input)
		if actual != expected {
			t.Errorf("For input %s, expected result <%s> but got instead <%s>", input, expected, actual)
		}

	}

}
