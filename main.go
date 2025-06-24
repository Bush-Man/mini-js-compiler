package main

import (
	"fmt"

	"github.com/Bush-Man/mini-js-compiler/lexer"
)
func main(){

	// var tokens []interface{}
	input := "123.4"
	lexer := lexer.NewLexer(input)
	for range input {

		token := lexer.NextToken()
		fmt.Println(*token)  // Implement proper String() for Token if needed
		if token.Kind == "EOF" {
			break
		}
	}

}