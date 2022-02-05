package lexer

import "gomonkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0 // zero is the ASCII code for the "NUL", lexer reached end of input
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var currentToken token.Token
	switch lexer.ch {
	case '=':
		currentToken = newToken(token.ASSIGN, lexer.ch)
	case ';':
		currentToken = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		currentToken = newToken(token.LPAREN, lexer.ch)
	case ')':
		currentToken = newToken(token.RPAREN, lexer.ch)
	case ',':
		currentToken = newToken(token.COMMA, lexer.ch)
	case '+':
		currentToken = newToken(token.PLUS, lexer.ch)
	case '{':
		currentToken = newToken(token.LBRACE, lexer.ch)
	case '}':
		currentToken = newToken(token.RBRACE, lexer.ch)
	case 0:
		currentToken.Type = token.EOF
		currentToken.Literal = ""
	}
	lexer.readChar()
	return currentToken
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
