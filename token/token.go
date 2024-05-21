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

    LT = "<"

    // Delimiters
    SEMICOLON = ";"

    LBRACE = "{"
    RBRACE = "}"

    EQ     = "=="

    NOT   = "NOT"
    AND   = "AND"
    OR    = "OR"
    IF    = "IF"
    ELSE  = "ELSE"
    WHILE = "WHILE"
    FOR   = "FOR"
    TRUE  = "TRUE"
    FALSE = "FALSE"
)

// Every other token is the same in both languages
var whileToGoTokens = map[TokenType]TokenType{
    WHILE:  FOR,
    AND:    AND,
    NOT:    NOT, 
}

var keywords = map[string]TokenType{
    "while":  WHILE,
    "if":     IF, 
    "else":   ELSE,
    "not":    NOT,
    "for":    FOR,
    "true":   TRUE,
    "or":     OR,
    "and":    AND,
    "false":  FALSE,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENT
}
