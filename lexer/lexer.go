package lexer

import (
	"github.com/negarciacamilo/go_interpreter/token"
)

type Lexer interface {
	NextToken() token.Token
}

type lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL"
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *lexer) readIdentifier() string {
	// Initial index
	position := l.position
	for isLetter(l.ch) {
		// Read char will fetch the last position in the identifier
		l.readChar()
	}
	return l.input[position:l.position]
}

// TODO: Duplicated code here, might want to join with readIdentifier
func (l *lexer) readNumber() string {
	// Initial index
	position := l.position
	for isDigit(l.ch) {
		// Read char will fetch the last position in the identifier
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		// Since ch is a rune, we should look ahead if the next char is a == so we can create the EQ operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return t
}

func (l *lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// This function will help lookin ahead of the char so we can branch the switch statement (i.e we can't use == since ch is a rune)
func (l *lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
