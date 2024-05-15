package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + Literals
    IDENT = "IDENT"
    INT   = "INT"

    // Operators 
    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    ASTERISK = "*" 
    BANG     = "!"

    LT = "<"
    GT = ">"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LBRACE = "{"
    RBRACE = "}"

    EQ     = "=="
    NOT_EQ = "NOT"

    AND = "AND"
    OR  = "OR"

    AMPERSAND = "&&"
    BAR       = "||"
    BANG_EQ   = "!="

    // -- Might not need the parentheses --
    LPAREN = "("
    RPAREN = ")"

    IF    = "IF"
    ELSE  = "ELSE"
    WHILE = "WHILE"
    FOR   = "FOR"
)

// Every other toekn is the same in both languages
var whileToGoTokens = map[TokenType]TokenType{
    WHILE:  FOR,
    AND:    AMPERSAND,
    OR:     BAR,
    NOT_EQ: BANG_EQ, 
}

var keywords = map[string]TokenType{
    "while":  WHILE,
    "if":     IF, 
    "else":   ELSE,
    "or":     OR,
    "and":    AND,
    "not":    NOT_EQ,
    "for":    FOR,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENT
}
