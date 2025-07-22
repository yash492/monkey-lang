package lexer

import (
	"monkey/token"
	"testing"
)

func TestLexer(t *testing.T) {
	type testCase struct {
		input  string
		output []token.Token
	}

	tests := []testCase{
		{
			input: `let x = 5;
					"foobar";	
			`,
			output: []token.Token{
				newToken(token.LET, "let"),
				newToken(token.IDENT, "x"),
				newToken(token.ASSIGN, "="),
				newToken(token.INT, "5"),
				newToken(token.SEMICOLON, ";"),
				newToken(token.STRING, "foobar"),
				newToken(token.SEMICOLON, ";"),

			},
		},
		{
			input: "(x + y) * 100 - (50 / 70) ",
			output: []token.Token{
				newToken(token.LPAREN, "("),
				newToken(token.IDENT, "x"),
				newToken(token.PLUS, "+"),
				newToken(token.IDENT, "y"),
				newToken(token.RPAREN, ")"),
				newToken(token.ASTERISK, "*"),
				newToken(token.INT, "100"),
				newToken(token.MINUS, "-"),
				newToken(token.LPAREN, "("),
				newToken(token.INT, "50"),
				newToken(token.SLASH, "/"),
				newToken(token.INT, "70"),
				newToken(token.RPAREN, ")"),
			},
		},
		{
			input: "if (x == 10) { return false } else { return 6 } fn(x) { return true } $ %",
			output: []token.Token{
				newToken(token.IF, "if"),
				newToken(token.LPAREN, "("),
				newToken(token.IDENT, "x"),
				newToken(token.EQ, "=="),
				newToken(token.INT, "10"),
				newToken(token.RPAREN, ")"),
				newToken(token.LBRACE, "{"),
				newToken(token.RETURN, "return"),
				newToken(token.FALSE, "false"),
				newToken(token.RBRACE, "}"),
				newToken(token.ELSE, "else"),
				newToken(token.LBRACE, "{"),
				newToken(token.RETURN, "return"),
				newToken(token.INT, "6"),
				newToken(token.RBRACE, "}"),
				newToken(token.FUNCTION, "fn"),
				newToken(token.LPAREN, "("),
				newToken(token.IDENT, "x"),
				newToken(token.RPAREN, ")"),
				newToken(token.LBRACE, "{"),
				newToken(token.RETURN, "return"),
				newToken(token.TRUE, "true"),
				newToken(token.RBRACE, "}"),
				newToken(token.ILLEGAL, "$"),
				newToken(token.ILLEGAL, "%"),
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		i := 0
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if i >= len(tt.output) {
				t.Fatalf("wrong token generated length. generated length - %v output length %v", i, len(tt.output))

			}
			if tok.Literal != tt.output[i].Literal {
				t.Fatalf("wrong literal. original %s & expected %s", tok.Literal, tt.output[i].Literal)
			}
			if tok.Type != tt.output[i].Type {
				t.Fatalf("wrong type. original %s & expected %s", tok.Type, tt.output[i].Type)
			}
			i++
		}

		if i != len(tt.output) {
			t.Fatalf("test %q: Incorrect number of tokens. got %d, expected %d", tt.input, i, len(tt.output))
		}
	}
}

func newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
