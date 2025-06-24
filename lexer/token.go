package lexer

type TokenKind string

const (

	// Arithmetic Operators
	PLUS              = "+"
	MINUS             = "-"
	ASTERISK          = "*"
	SLASH             = "/"
	PERCENT           = "%"
	INCREMENT         = "++"
	EXPONENTIAL       = "**"
	DECREAMENT       = "--"
	
	// Assignment Operators
	EQUALS            = "="
	PLUS_EQUALS       = "+="
	TIMES_EQUALS      = "*="
	DIVIDE_EQUALS     = "/="
	MINUS_EQUALS      = "-="
	MOD_EQUALS        = "%=" 
	STAR_EQUALS       = "*="  
	SLASH_EQUALS      = "/="  
	PERCENT_EQUALS    = "%="  
	LEFT_SHIFT_EQUALS = "<<="  
	RIGHT_SHIFT_EQUALS = ">>="  
	BITWISE_AND_EQUALS = "&="  
	BITWISE_OR_EQUALS  = "|="  
	BITWISE_XOR_EQUALS = "^="  


	//Comparison Operators
	STRICT_EQUALS   = "==="
	LOOSE_EQUALS    = "=="
	NOT_EQUALS        = "!="
	STRICT_NOT_EQUALS = "!=="
	LESS_EQUALS       = "<="
	GREATER_EQUALS    = ">="
	LESS_THAN         = "<"
	GREATER_THAN      = ">"
	
	// Logical Operators
	NOT               = "!"
	AND               = "&&"
	OR                = "||"
	BANG              = "!"
	NULLISH_COALESCING = "??"  
	OPTIONAL_CHAINING = "?." 


	// Bitwise Operators
	AMPERSAND         = "&"
	PIPE              = "|"
	CARET             = "^"  
	TILDE             = "~"  
	LEFT_SHIFT        = "<<"  
	RIGHT_SHIFT       = ">>"  
	UNSIGNED_RIGHT_SHIFT = ">>>"  


	// Punctuations
	COMMA             = ","  
	DOT               = "."  
	COLON             = ":"  
	SEMICOLON         = ";"  
	QUESTION          = "?"  
	AT                = "@"  
	HASH              = "#"  
	DOLLAR            = "$"  
	UNDERSCORE        = "_"
	BACKTICK          = "`"
	QUOTE             = "'"  
	DOUBLE_QUOTE      = "\"\""
	BACKSLASH         = "\\"  
	LEFT_PAREN        = "("  
	RIGHT_PAREN       = ")"  
	LEFT_BRACKET      = "["  
	RIGHT_BRACKET     = "]"  
	LEFT_BRACE        = "{"  
	RIGHT_BRACE       = "}"  
	ARROW             = "=>"  
	SPREAD            = "..."  


	//Literals
	NUMBER ="NUMBER"

	// Special
	EOF     = "EOF"
	UNKNOWN = "UNKNOWN"
	NO_TOKENS = "NO_TOKENS"
)

type TextSpan struct {
	Start   int
	End     int
	Literal string
}

func NewTextSpan(start int, end int, literal string) *TextSpan {
	return &TextSpan{
		Start:   start,
		End:     end,
		Literal: literal,
	}
}

type Token struct {
	Kind TokenKind
	Span TextSpan
}

func NewToken(kind TokenKind, span TextSpan) Token {
	return Token{
		Kind: kind,
		Span: span,
	}
}
