package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	currPosition int // currPosition is the current position of the char
	nextPosition int // nextPosition is used to query the char and store it in the char field and is incremented by 1
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		currPosition: 0,
		nextPosition: 0,
		char:         0,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {

	l.skipWhitespace()

	var tok token.Token
	switch l.char {
	case '=':
		nextChar := l.peekChar()
		if nextChar == '=' {
			tok.Literal = string(nextChar) + string(l.char)
			tok.Type = token.EQ
			l.readChar()
		} else {
			tok.Literal = string(l.char)
			tok.Type = token.ASSIGN
		}
	case '!':
		nextChar := l.peekChar()
		if nextChar == '=' {
			tok.Literal = string(l.char) + string(nextChar)
			tok.Type = token.NOTEQ
			l.readChar()
		} else {
			tok.Literal = string(l.char)
			tok.Type = token.BANG
		}
	case ';':
		tok.Literal = string(l.char)
		tok.Type = token.SEMICOLON

	case '>':
		tok.Literal = string(l.char)
		tok.Type = token.GT

	case '<':
		tok.Literal = string(l.char)
		tok.Type = token.LT

	case '+':
		tok.Literal = string(l.char)
		tok.Type = token.PLUS

	case '-':
		tok.Literal = string(l.char)
		tok.Type = token.MINUS

	case '*':
		tok.Literal = string(l.char)
		tok.Type = token.ASTERISK

	case '/':
		tok.Literal = string(l.char)
		tok.Type = token.SLASH

	case ',':
		tok.Literal = string(l.char)
		tok.Type = token.COMMA

	case '(':
		tok.Literal = string(l.char)
		tok.Type = token.LPAREN

	case ')':
		tok.Literal = string(l.char)
		tok.Type = token.RPAREN

	case '{':
		tok.Literal = string(l.char)
		tok.Type = token.LBRACE

	case '}':
		tok.Literal = string(l.char)
		tok.Type = token.RBRACE

	case 0:
		tok.Literal = string(l.char)
		tok.Type = token.EOF

	default:
		if isLetter(l.char) {
			identifier := l.getTextEntity(isLetter)
			tokenType := token.LookupIdent(identifier)
			tok.Type = tokenType
			tok.Literal = identifier
			return tok

		} else if isNumber(l.char) {
			number := l.getTextEntity(isNumber)
			tok.Type = token.INT
			tok.Literal = number
			return tok
		} else {
			tok.Literal = string(l.char)
			tok.Type = token.ILLEGAL
		}

	}

	l.readChar()
	return tok

}

// skipWhitespace removes all sorts of whitespaces such as spaces, new lines, tabs and carriage return
func (l *Lexer) skipWhitespace() {
	for l.char == '\n' || l.char == '\t' || l.char == '\r' || l.char == ' ' {
		l.readChar()
	}

}

// readChar reads a character and advances Lexer positions.
// Reads Lexer.char, increments Lexer.currPosition, and Lexer.nextPosition.
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

// peekChar peeks the next char from the input
func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}

func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch == '_' {
		return true
	}
	return false
}

func isNumber(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

// getTextEntity fetches a text entity from the input. It consumes characters
// as long as the provided function 'fn' returns true for each character.
//
// Parameters:
//   - fn: A function (byte -> bool) that determines which characters to include.
func (l *Lexer) getTextEntity(fn func(ch byte) bool) string {
	startPos := l.currPosition
	for fn(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.currPosition]

}
