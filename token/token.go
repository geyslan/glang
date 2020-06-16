package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOL     = "EOL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// Enclosures
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"

	// Types
	TYPEINTAUTO  = "INTAUTO"
	TYPEINTALLOC = "INTALLOC"
	TYPEINTFIXED = "INTFIXED"
	TYPEINTVOLAT = "INTVOLAT"
)

var keywords = map[string]TokenType{
	"func": FUNCTION,
	"var":  VAR,
}

var types = map[string]TokenType{
	"int:auto":  TYPEINTAUTO,
	"int:alloc": TYPEINTALLOC,
	"int:fixed": TYPEINTFIXED,
	"int:volat": TYPEINTVOLAT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if tok, ok := types[ident]; ok {
		return tok
	}
	return IDENT
}
