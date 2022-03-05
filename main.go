package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	re := regexp.MustCompile(`^[ 　\t]+`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := re.ReplaceAllStringFunc(scanner.Text(), func(s string) string {
			spaceCount := utf8.RuneCountInString(s)
			return strings.Repeat(" ", spaceCount)
		})
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
