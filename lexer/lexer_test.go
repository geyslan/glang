package lexer

import (
	"testing"

	"glang/token"
)

func TestNextToken(t *testing.T) {
	input := `var vauto int:auto = 10
	
	var (
		first_fixed1 int:fixed
		second_auto2 int:alloc
	)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "vauto"},
		{token.TYPEINTAUTO, "int:auto"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.VAR, "var"},
		{token.LPAREN, "("},
		{token.IDENT, "first_fixed1"},
		{token.TYPEINTFIXED, "int:fixed"},
		{token.IDENT, "second_auto2"},
		{token.TYPEINTALLOC, "int:alloc"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

	}
}
