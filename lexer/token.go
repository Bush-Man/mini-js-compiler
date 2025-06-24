package lexer

type TokenKind string

const (

	// Operators
	PLUS              = "+"
	MINUS             = "-"
	ASTERISK          = "*"
	ASSIGN            = "="
	SLASH             = "/"
	BANG              = "!"
	MOD               = "%"
	LESS              = "<"
	GREATER           = ">"
	EQUALS            = "=="
	STRICT_EQUALS     = "==="
	NOT_EQUALS        = "!="
	NOT_STRICT_EQUALS = "!=="
	PLUS_EQUALS       = "+="
	MINUS_EQUALS      = "-="
	TIMES_EQUALS      = "*="
	DIVIDE_EQUALS     = "/="
	MOD_EQUALS        = "%="
	PLUS_PLUS         = "++"
	MINUS_MINUS       = "--"
	AMPERSAND         = "&"
	PIPE              = "|"
	AND               = "&&"
	OR                = "||"

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
