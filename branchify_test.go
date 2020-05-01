package main

import (
	"testing"
)

func TestBranchifyBasic(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{args: []string{}, expected: ""},
		{args: []string{"one"}, expected: "one"},
		{args: []string{"one", "two"}, expected: "one-two"},
		{args: []string{"one two"}, expected: "one-two"},
		{args: []string{"ONE"}, expected: "one"},
		{args: []string{"One", "Two"}, expected: "one-two"},
	}
	for _, tt := range tests {
		got := branchify(tt.args, "-", "", "")
		if got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}

func TestBranchifyWithNamespace(t *testing.T) {
	var tests = []struct {
		args     []string
		ns       string
		expected string
	}{
		{args: []string{}, ns: "t", expected: "t@"},
		{args: []string{"one"}, ns: "t", expected: "t@one"},
		{args: []string{"one", "two"}, ns: "t", expected: "t@one-two"},
	}
	for _, tt := range tests {
		got := branchify(tt.args, "-", tt.ns, "@")
		if got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}

func TestBranchifyIllegalChars(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{args: []string{"one:"}, expected: "one"},
		{args: []string{"one?????"}, expected: "one"},
	}
	for _, tt := range tests {
		got := branchify(tt.args, "-", "", "")
		if got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}
