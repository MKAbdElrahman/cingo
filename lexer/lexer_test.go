package lexer_test

import (
	"cingo/lexer"
	"regexp"
	"testing"
)

func TestLexSemicolon(t *testing.T) {

	var regexMap = map[string]*regexp.Regexp{
		"Semicolon": regexp.MustCompile(`^;`),
	}

	input := " ; ; "
	expectedTokens := []string{";", ";"}
	tokens, err := lexer.Lex(input, regexMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareSlices(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, tokens)
	}
}

func TestLexbrace(t *testing.T) {

	var regexMap = map[string]*regexp.Regexp{
		"Semicolon":         regexp.MustCompile(`^;`),
		"Openbrace":         regexp.MustCompile(`^{`),
		"Closedbrace":       regexp.MustCompile(`^}`),
		"Openparenthesis":   regexp.MustCompile(`^\(`),
		"Closedparenthesis": regexp.MustCompile(`^\)`),
	}

	input := " { ; } ()"
	expectedTokens := []string{"{", ";", "}","(",")"}
	tokens, err := lexer.Lex(input, regexMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !compareSlices(tokens, expectedTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, tokens)
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
