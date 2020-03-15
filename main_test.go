package main

import (
	"fmt"
	"testing"
)

// func TestFirstVerse(t *testing.T) {
// 	got := Verse()
// 	expected := "99 bottles of beer on the wall, 99 bottles of beer.\n" +
// 		"Take one down, pass it around. 98 bottles of beer on the wall."

// 	if got != expected {
// 		t.Errorf("Got: %s\nExp: %s", got, expected)
// 	}
// }

func TestVerse(t *testing.T) {
	tests := []struct {
		bottles  int
		expected string
	}{
		{
			bottles: 99,
			expected: "99 bottles of beer on the wall, 99 bottles of beer.\n" +
				"Take one down, pass it around. 98 bottles of beer on the wall.",
		},
		{
			bottles: 98,
			expected: "98 bottles of beer on the wall, 98 bottles of beer.\n" +
				"Take one down, pass it around. 97 bottles of beer on the wall.",
		},
		{
			bottles: 2,
			expected: "2 bottles of beer on the wall, 2 bottles of beer.\n" +
				"Take one down, pass it around. 1 bottle of beer on the wall.",
		},
		{
			bottles: 1,
			expected: "1 bottle of beer on the wall, 1 bottle of beer.\n" +
				"Take it down, pass it around. No more bottles of beer on the wall.",
		},
		{
			bottles: 0,
			expected: "No more bottles of beer on the wall, no more bottles of beer.\n" +
				"Go to the store, buy some more. 99 bottles of beer on the wall.",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.bottles), func(t *testing.T) {
			got := GetVerse(tt.bottles).String()
			expected := tt.expected
			if got != expected {
				t.Errorf("Got:\n%s\nExpected:\n%s", got, expected)
			}
		})
	}
}
