// Package lexer provides the lexer for Monkey
package lexer

import "github.com/yngtodd/monkey/token"

// A Lexer holds our text input, positions, and current character.
type Lexer struct {
    input        string
    position     int  // current position in input (points to current char)
    readPosition int  // current reading position (after current char)
    ch           byte // current char under consideration
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

// readChar reads the next character in our input 
// and increments stores the next position to be read.
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
        case '=':
	    tok = newToken(token.ASSIGN, l.ch)
	case ';':
            tok = newToken(token.SEMICOLON, l.ch)
	case '(':
	    tok = newToken(token.LPAREN, l.ch)
	case ')':
	    tok = newToken(token.RPAREN, l.ch)
	case ',':
	    tok = newToken(token.COMMA, l.ch)
	case '+':
	    tok = newToken(token.PLUS, l.ch)
	case '{':
	    tok = newToken(token.LBRACE, l.ch)
	case '}':
	    tok = newToken(token.RBRACE, l.ch)
	case 0:
	    tok.Literal = ""
	    tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
