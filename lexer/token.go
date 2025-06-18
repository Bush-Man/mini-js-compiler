package lexer;

type TokenKind string

const (


	// Operators
    PLUS="+"
	MINUS="-"
	ASTERISK="*"
    ASSIGN="="
	SLASH="/"
	BANG="!"
	MOD="%"
	LESS="<"
	GREATER=">"
	EQUALS="=="
	STRICT_EQUALS="==="
	NOT_EQUALS="!="
	NOT_STRICT_EQUALS="!=="
	PLUS_EQUALS="+="
	MINUS_EQUALS="-="
	TIMES_EQUALS="*="
	DIVIDE_EQUALS="/="
	MOD_EQUALS="%="
	PLUS_PLUS="++"
	MINUS_MINUS="--"
	AMPERSAND="&"
	PIPE="|"
	AND="&&"
	OR="||"
	

	// Special
	EOF="EOF"
	UNKNOWN ="UNKNOWN"


)