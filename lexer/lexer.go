package lexer

import (
	"monkeylang/token"
)

var keywords = map[string]token.TokenType{
	"let":    token.LET,
	"fn":     token.FUNCTION,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
}

func lookupIdent(ident string) token.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return token.IDENT
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

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
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Value: "="}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Value: ";"}
	case '(':
		tok = token.Token{Type: token.LPAREN, Value: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Value: ")"}
	case ',':
		tok = token.Token{Type: token.COMMA, Value: ","}
	case '+':
		tok = token.Token{Type: token.PLUS, Value: "+"}
	case '-':
		tok = token.Token{Type: token.MINUS, Value: "-"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Value: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Value: "}"}
	case '!':
		tok = token.Token{Type: token.BANG, Value: "!"}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Value: "*"}
	case '/':
		tok = token.Token{Type: token.SLASH, Value: "/"}
	case '<':
		tok = token.Token{Type: token.LESS_THAN_OR_EQUAL, Value: "<"}
	case '>':
		tok = token.Token{Type: token.GREATER_THAN_OR_EQUAL, Value: ">"}
	case 0:
		tok = token.Token{Type: token.EOF, Value: ""}
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = lookupIdent(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok.Value = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Value: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// TODO: support floats
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
