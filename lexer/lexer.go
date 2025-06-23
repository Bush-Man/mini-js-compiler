package lexer

import (
	"fmt"
	"unicode"
)

   
type Lexer struct{
  input string
  currentPosition int

}

func NewLexer(input string) *Lexer {
  lexer := &Lexer{
    input: input,
    currentPosition:0,
   
   }
    return lexer
}

func (lexer *Lexer) NextToken()*Token{
  var currentToken *Token = nil
   print(currentToken)

  if lexer.currentPosition >= len(lexer.input){
    eofChar := "\x00"
    token := NewToken(EOF, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, eofChar))
    currentToken = &token
    print("EOF")
  }
 

  
  for lexer.currentPosition < len(lexer.input){
     currentChar := lexer.input[lexer.currentPosition]

   
    if lexer.isNumber(currentChar) {
      numberString := lexer.consumeNumber()
      token := NewToken(INTERGER,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+1,string(numberString)))
      print("NUMBER")
      currentToken = &token

    }else{
      lexer.currentPosition++
    }
   }
  return currentToken

}

func (lexer *Lexer) consumePunctuation()*Token{
  switch lexer.input[lexer.currentPosition] {
  case '+':
     token := NewToken(PLUS, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "+"))
     lexer.currentPosition++
     return &token
  case '-':
     token := NewToken(MINUS, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "-"))
     lexer.currentPosition++
     return &token    
  case '*':
      token := NewToken(ASTERISK, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "*"))
      lexer.currentPosition++
      return &token
  case '/':
       token := NewToken(SLASH, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "/"))
       lexer.currentPosition++
       return &token 
         
  default:
     token := NewToken(UNKNOWN, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, string(lexer.input[lexer.currentPosition])))
     lexer.currentPosition++
     return &token
  }

}

func (lexer *Lexer)consumeNumber()string{
  fullNumber := ""
  currentChar := lexer.input[lexer.currentPosition]
  for lexer.currentPosition < len(lexer.input) && lexer.isNumber(currentChar){
    fullNumber+=string(currentChar)
    lexer.currentPosition++
  }
  return fullNumber
}
func (lexer *Lexer) isNumber(char byte)bool{
  fmt.Printf("NUmber: %c\n", char)
  return unicode.IsDigit(rune(char))
}

func (lexer *Lexer) isLetter(char byte)bool{
  return unicode.IsLetter(rune(char))
}
