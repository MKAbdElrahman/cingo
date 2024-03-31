package main

import (
	"cingo/lexer"
	"fmt"

	"github.com/charmbracelet/log"
)

func main() {
	input := `
	int	main	(	void)	{	return	0	;	}
	`
	tokens, err := lexer.Lex(input, lexer.DefaultRegexMap)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("+----------+--------------------+")
	fmt.Println("|     Value     |        Type       |")
	fmt.Println("+----------+--------------------+")
	for _, token := range tokens {
		fmt.Printf("| %-15v | %-20v |\n", token.Value, token.Type)
	}
}
