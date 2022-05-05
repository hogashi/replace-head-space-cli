package main

import "testing"

func TestReplaceHead(t *testing.T) {
	cases := []struct {
		target, replaceWith, in, want string
	}{
		{`[ 　\t]+`, " ", "no space", "no space"},
		{`[ 　\t]+`, " ", " 1 half", " 1 half"},
		{`[ 　\t]+`, " ", "　1 full", " 1 full"},
		{`[ 　\t]+`, " ", "	1 tab", " 1 tab"},
		{`[ 　\t]+`, " ", " 　	1 half 1 full 1 tab (3)", "   1 half 1 full 1 tab (3)"},
		{`[ 　\t]+`, "ABC", "no space", "no space"},
		{`[ 　\t]+`, "ABC", " 1 half", "ABC1 half"},
		{`[ 　\t]+`, "ABC", "　1 full", "ABC1 full"},
		{`[ 　\t]+`, "ABC", "	1 tab", "ABC1 tab"},
		{`[ 　\t]+`, "ABC", " 　	1 half 1 full 1 tab (3)", "ABCABCABC1 half 1 full 1 tab (3)"},
		{`[a-z]+`, "ABC", "no space", "ABCABC space"},
		{`[a-z]+`, "ABC", "	1 tab", "	1 tab"},
	}
	for _, c := range cases {
		got := ReplaceHead(c.in, c.target, c.replaceWith)
		if got != c.want {
			t.Errorf("ReplaceHead(%q, %q, %q) == %q, want %q", c.in, c.target, c.replaceWith, got, c.want)
		}
	}
}
