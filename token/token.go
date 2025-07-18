package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN                = "="
	PLUS                  = "+"
	MINUS                 = "-"
	ASTERISK              = "*"
	SLASH                 = "/"
	EQUALS                = "=="
	BANG                  = "!"
	NOT_EQUALS            = "!="
	GREATER               = ">"
	LESS                  = "<"
	GREATER_THAN_OR_EQUAL = ">="
	LESS_THAN_OR_EQUAL    = "<="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)
