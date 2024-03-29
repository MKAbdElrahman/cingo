package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

var DefaultRegexMap = map[string]*regexp.Regexp{
	"Semicolon": regexp.MustCompile(`^;`),
}

func Lex(input string, regexMap map[string]*regexp.Regexp) ([]string, error) {
	var tokens []string
	for len(input) != 0 {
		input = strings.TrimSpace(input)
		longestMatch := findLongestMatch(input, regexMap)
		if len(longestMatch) > 0 {
			tokens = append(tokens, longestMatch)
			input = strings.TrimPrefix(input, longestMatch)
			continue
		} else {
			return []string{}, fmt.Errorf("failed to parse %q", input)
		}
	}
	return tokens, nil
}

func findLongestMatch(s string, regexMap map[string]*regexp.Regexp) string {
	maxLength := 0
	longestMatch := ""
	for _, v := range regexMap {
		match := v.FindString(s)
		if match == "" {
			continue
		}
		if len(match) > maxLength {
			maxLength = len(match)
			longestMatch = match
		}
	}
	return longestMatch
}
