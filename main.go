package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/jessevdk/go-flags"
)

var defaultTarget = `^[ 　\t]+`
var defaultReplaceWith = " "

func main() {
	var opts struct {
		Target      string `short:"t" long:"target" description:"replace target RegExp (default: '^[ 　\t]+')"`
		ReplaceWith string `short:"w" long:"with" description:"replace with this string (default: ' ')"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	target := defaultTarget
	if opts.Target != "" {
		target = opts.Target
	}
	re := regexp.MustCompile(target)

	replaceWith := defaultReplaceWith
	if opts.ReplaceWith != "" {
		replaceWith = opts.ReplaceWith
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := re.ReplaceAllStringFunc(scanner.Text(), func(s string) string {
			spaceCount := utf8.RuneCountInString(s)
			return strings.Repeat(replaceWith, spaceCount)
		})
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
