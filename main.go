package main

import (
	"fmt"

	"github.com/Bush-Man/mini-js-compiler/lexer"
)
func main(){

	// var tokens []interface{}
	input := "1234"
	lexer := lexer.NewLexer(input)
	for{

		token := lexer.NextToken()
		fmt.Println(*token)  // Implement proper String() for Token if needed
		if token.Kind == "EOF" {
			break
		}
	}

}