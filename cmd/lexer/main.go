package main

import (
	"cingo/lexer"
	"fmt"
	"regexp"

	"github.com/charmbracelet/log"
)

var regexMap = map[string]*regexp.Regexp{
	"Semicolon": regexp.MustCompile(`^;`),
}

func main() {
	input := ""
	tokens, err := lexer.Lex(input, regexMap)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(tokens)
}
