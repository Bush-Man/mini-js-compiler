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
      numberString := lexer.tokenizeNumber()
      end := lexer.currentPosition
      token := NewToken(NUMBER,*NewTextSpan(start,end,numberString))
      return &token

    }

   
  return lexer.tokenizeArithmeticOperators()

}
// Arithmetic operators tokenizer
func (lexer *Lexer) tokenizeArithmeticOperators()*Token{
  switch lexer.input[lexer.currentPosition] {
  case '+':
      if lexer.currentPosition+1 < len(lexer.input) && lexer.input[lexer.currentPosition+1]== '+'{
          token := NewToken(INCREMENT, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+2, "++"))
          lexer.currentPosition += 2
          return &token
      }else if  lexer.currentPosition + 1 < len(lexer.input) && lexer.input[lexer.currentPosition +1]=='='{
           token := NewToken(PLUS_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
           lexer.currentPosition+=2
           return &token
           
      }else{

        token := NewToken(PLUS, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "+"))
        lexer.currentPosition++
        return &token
      }
  case '-':
    if lexer.currentPosition+1 < len(lexer.input) && lexer.input[lexer.currentPosition+1] == '-'{
      token := NewToken(DECREAMENT,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2, "--"))
      lexer.currentPosition +=2
      return &token
      }else if  lexer.currentPosition + 1 < len(lexer.input) && lexer.input[lexer.currentPosition +1]=='='{
        token := NewToken(MINUS_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
        lexer.currentPosition+=2
        return &token
    }else{

      token := NewToken(MINUS, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "-"))
      lexer.currentPosition++
      return &token    
    }
  case '*':
     if lexer.currentPosition+1 < len(lexer.input) && lexer.input[lexer.currentPosition+1] == '*' {
      token := NewToken(EXPONENTIAL,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2, "**"))
      lexer.currentPosition +=2
      return &token
      }else if  lexer.currentPosition + 1 < len(lexer.input) && lexer.input[lexer.currentPosition +1]=='='{
        token := NewToken(TIMES_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
        lexer.currentPosition+=2
        return &token
     }else{

       token := NewToken(ASTERISK, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, "*"))
       lexer.currentPosition++
       return &token
      }
  case '/':
       if lexer.currentPosition +1 < len(lexer.input) && lexer.input[lexer.currentPosition+1]== '='{
        token := NewToken(SLASH_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
        lexer.currentPosition+=2
        return &token
       }else{

       
       token := NewToken(SLASH, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1,  string(lexer.input[lexer.currentPosition])))
       lexer.currentPosition++
       return &token 
       }
  case '%':
    if lexer.currentPosition +1 < len(lexer.input) && lexer.input[lexer.currentPosition+1]== '='{
      token := NewToken(MOD_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
      lexer.currentPosition+=2
      return &token
     }else{

     
     token := NewToken(SLASH, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, string(lexer.input[lexer.currentPosition])))
     lexer.currentPosition++
     return &token 
     }
       
         
  default:
     token := NewToken(UNKNOWN, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, string(lexer.input[lexer.currentPosition])))
     lexer.currentPosition++
     return &token
  }

}

// Assignment and Comparison Operators  Tokenizer

func (lexer *Lexer) tokenizeAssignmentAndComparisonOperators() *Token{

  switch lexer.input[lexer.currentPosition]{
  case '=':
    if lexer.currentPosition+1 < len(lexer.input) && lexer.input[lexer.currentPosition+1] == '=' && lexer.input[lexer.currentPosition+2] != '=' {
      token := NewToken(LOOSE_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
      lexer.currentPosition+=2
      return &token
    }else if lexer.currentPosition+2 < len(lexer.input) && lexer.input[lexer.currentPosition+1] == '=' && lexer.input[lexer.currentPosition+2]=='='{
      token := NewToken(STRICT_EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+3,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])+string(lexer.input[lexer.currentPosition+2])))
      lexer.currentPosition+=3
      return &token
    }else if lexer.currentPosition+1 < len(lexer.input) && lexer.input[lexer.currentPosition+1]=='>'{
       token := NewToken(ARROW,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+2,string(lexer.input[lexer.currentPosition])+string(lexer.input[lexer.currentPosition+1])))
       lexer.currentPosition+=2
       return &token
    
    }else{
      token := NewToken(EQUALS,*NewTextSpan(lexer.currentPosition,lexer.currentPosition+1,string(lexer.input[lexer.currentPosition])))
      lexer.currentPosition++
      return &token
    }

  default:
    token := NewToken(UNKNOWN, *NewTextSpan(lexer.currentPosition, lexer.currentPosition+1, string(lexer.input[lexer.currentPosition])))
    lexer.currentPosition++
    return &token   
  }

}


func (lexer *Lexer) tokenizeNumber() string {
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
