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
		args      []string
		namespace string
		expected  string
	}{
		{args: []string{}, namespace: "t", expected: "t:"},
		{args: []string{"one"}, namespace: "t", expected: "t:one"},
		{args: []string{"one", "two"}, namespace: "t", expected: "t:one-two"},
	}
	for _, tt := range tests {
		got := branchify(tt.args, "-", tt.namespace, ":")
		if got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}
