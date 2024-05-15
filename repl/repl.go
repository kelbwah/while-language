package repl

import (
    "bufio"
    "fmt"
    "io"
    "while/lexer"
    // "while/token"
    "while/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()

        if (!scanned) {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)
        p := parser.New(l)

        statements := p.ParseProgram()

        for _, statement := range statements.Statements {
            fmt.Println(statement) 
        }

        /*
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
        */
    }
}
