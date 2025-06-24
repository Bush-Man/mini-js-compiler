package lexer

import (
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
 

  if lexer.currentPosition >= len(lexer.input){
    eofChar := "\x00"
    token := NewToken(EOF, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, eofChar))
    return &token
  }
 

  
 
     currentChar := lexer.input[lexer.currentPosition]

   
    if lexer.isNumber(currentChar) {
      start := lexer.currentPosition
      numberString := lexer.consumeNumber()
      end := lexer.currentPosition
      token := NewToken(NUMBER,*NewTextSpan(start,end,numberString))
      return &token

    }

   
  return lexer.consumePunctuation()

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

func (lexer *Lexer) consumeNumber() string {
  fullNumber := ""
  for lexer.currentPosition < len(lexer.input) &&
      (lexer.isNumber(lexer.input[lexer.currentPosition]) || lexer.input[lexer.currentPosition] == '.') {
      fullNumber += string(lexer.input[lexer.currentPosition])
      lexer.currentPosition++
  }
  return fullNumber
}
func (lexer *Lexer) isNumber(char byte)bool{
  return unicode.IsDigit(rune(char))
}

func (lexer *Lexer) isLetter(char byte)bool{
  return unicode.IsLetter(rune(char))
}
