package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

type MatchPair struct {
	Type  string
	Value string
}

var DefaultRegexMap = map[string]*regexp.Regexp{
	"Semicolon":         regexp.MustCompile(`^;`),
	"Openbrace":         regexp.MustCompile(`^{`),
	"Closedbrace":       regexp.MustCompile(`^}`),
	"Openparenthesis":   regexp.MustCompile(`^\(`),
	"Closedparenthesis": regexp.MustCompile(`^\)`),
	"Constant":          regexp.MustCompile(`^[0-9]+\b`),
	"Identifier":        regexp.MustCompile(`^[a-zA-Z_]\w*\b`),
}

// a keyword is a special type of identifier
func convertTypeIfKeyword(pair *MatchPair) MatchPair {
	switch pair.Value {
	case "int":
		pair.Type = "Int Keyword"
		return *pair
	case "void":
		pair.Type = "Void Keyword"
		return *pair
	case "return":
		pair.Type = "Return Keyword"
		return *pair
	default:
		return *pair
	}
}

func Lex(input string, regexMap map[string]*regexp.Regexp) ([]MatchPair, error) {
	var tokens []MatchPair
	for len(input) != 0 {
		input = strings.TrimSpace(input)
		match := findLongestMatch(input, regexMap)
		if len(match.Value) > 0 {
			tokens = append(tokens, match)
			input = strings.TrimPrefix(input, match.Value)
			continue
		} else {
			return []MatchPair{}, fmt.Errorf("failed to parse %q", input)
		}
	}
	return tokens, nil
}

func findLongestMatch(s string, regexMap map[string]*regexp.Regexp) MatchPair {
	maxLength := 0
	longestMatch := ""
	longestMatchKey := ""
	for k, v := range regexMap {
		match := v.FindString(s)
		if match == "" {
			continue
		}
		if len(match) > maxLength {
			maxLength = len(match)
			longestMatch = match
			longestMatchKey = k
		}
	}
	return convertTypeIfKeyword(&MatchPair{
		Type:  longestMatchKey,
		Value: longestMatch,
	})
}
