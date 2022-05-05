package main

import "testing"

func TestReplaceHead(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"no space", "no space"},
		{" 1 half", " 1 half"},
		{"　1 full", " 1 full"},
		{"	1 tab", " 1 tab"},
		{"   3 half", "   3 half"},
		{"　　　3 full", "   3 full"},
		{"			3 tab", "   3 tab"},
		{" 　1 half, 1 full (2)", "  1 half, 1 full (2)"},
		{" 	1 half, 1 tab (2)", "  1 half, 1 tab (2)"},
		{"　	1 full, 1 tab (2)", "  1 full, 1 tab (2)"},
		{" 　	1 half, 1 full, 1 tab (3)", "   1 half, 1 full, 1 tab (3)"},
	}
	for _, c := range cases {
		got := ReplaceHead(c.in, `[ 　\t]+`, " ")
		if got != c.want {
			t.Errorf("ReplaceHead(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
