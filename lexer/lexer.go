/*
Package lexer implements a lexer, which takes source code
as input and outputs the tokens which represent the source code.

The lexer does not buffer or save tokens, it 
simply outputs the next token.
*/

package lexer

import "github.com/analyticalex/go-interpreter/token"

// TO DO: Support full Unicode range instead of only ASCII characters
type Lexer struct {
	input			string
	position		int // points to current char
	readPosition	int // points to next char, to "peek" ahead
	ch				byte // current char being examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Read current character, increments position and readPosition
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Returns token for current character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACK, l.ch)
	case ']':
		tok = newToken(token.RBRACK, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch)
	}
	
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
