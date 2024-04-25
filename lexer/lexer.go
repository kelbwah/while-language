package lexer

import (
    "while/token"
)

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}

    // l.readChar() <---- implement this

    return l
}

// Add in helper functions for moving tokens forward here
