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
		ReplaceWith string `short:"w" long:"with" description:"replace with this string"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`^[ 　\t]+`)

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
