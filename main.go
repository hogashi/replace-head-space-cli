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

var defaultTarget = `[ 　\t]+`
var defaultReplaceWith = " "

func main() {
	// parse options
	var opts struct {
		Target      string `short:"t" long:"target" description:"replace target RegExp (default: '[ 　\t]+')"`
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

	replaceWith := defaultReplaceWith
	if opts.ReplaceWith != "" {
		replaceWith = opts.ReplaceWith
	}

	// replace stdin, output to stdout
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputText := scanner.Text()
		result := ReplaceHead(inputText, target, replaceWith)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func ReplaceHead(text string, target string, replaceWith string) string {
	re := regexp.MustCompile("^" + target)
	return re.ReplaceAllStringFunc(text, func(s string) string {
		spaceCount := utf8.RuneCountInString(s)
		return strings.Repeat(replaceWith, spaceCount)
	})
}
