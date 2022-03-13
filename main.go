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

func main() {
	var opts struct {
		Target      string `short:"t" long:"target" description:"replace target RegExp"`
		ReplaceWith string `short:"w" long:"with" description:"replace with this string"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	target := `^[ 　\t]+`
	if opts.Target != "" {
		target = opts.Target
	}
	re := regexp.MustCompile(target)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := re.ReplaceAllStringFunc(scanner.Text(), func(s string) string {
			spaceCount := utf8.RuneCountInString(s)
			return strings.Repeat(opts.ReplaceWith, spaceCount)
		})
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
