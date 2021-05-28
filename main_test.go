package main

import (
	"testing"
)

// func TestReverseString(t *testing.T) {
// 	t.Run("reverse of Liam", func(t *testing.T) {
// 		str := "Liam"
// 		expected := "maiL"
// 		got := reverseString(str)
// 		if got != expected {
// 			t.Errorf("expected %s, got %s", expected, got)
// 		}
// 	})

// 	t.Run("reverse of foobar", func(t *testing.T) {
// 		str := "foobar"
// 		expected := "raboof"
// 		got := reverseString(str)
// 		if got != expected {
// 			t.Errorf("expected %s, got %s", expected, got)
// 		}
// 	})

// 	t.Run("reverse of testing", func(t *testing.T) {
// 		str := "testing"
// 		expected := "gnitset"
// 		got := reverseString(str)
// 		if got != expected {
// 			t.Errorf("expected %s, got %s", expected, got)
// 		}
// 	})
// }

func TestReverseString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "Liam", output: "maiL"},
		{input: "foobar", output: "raboof"},
		{input: "testing", output: "gnitset"},
	}

	for _, testCase := range tests {
		got := reverseString(testCase.input)
		expected := testCase.output
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	}
}
