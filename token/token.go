package token

type TokenType string

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers and literals
    IDENT = "IDENT"
    INT   = "INT"

    // Operators
    ASSIGN = "="
    PLUS   = "+"

    // Delimeters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
)

type Token struct {
    Type    TokenType
    Literal string
}
