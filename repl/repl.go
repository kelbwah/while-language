package repl

import (
    "bufio"
    "fmt"
    "io"
    "sync"
    "while/evaluator"
    "while/lexer"
    "while/object"
    "while/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, wg *sync.WaitGroup, inputChan <-chan string, latestOutput *string, mutex *sync.Mutex, doneChan chan struct{}) {
    defer wg.Done()
    fmt.Println("REPL Started")

    env := object.NewEnvironment()
    var replWg sync.WaitGroup

    // Goroutine to handle input from the terminal
    replWg.Add(1)
    go func() {
        defer replWg.Done()
        scanner := bufio.NewScanner(in)
        for {
            fmt.Print(PROMPT)
            if scanner.Scan() {
                line := scanner.Text()
                processLine(line, out, env, latestOutput, mutex, doneChan)
            } else {
                break
            }
        }
    }()

    // Goroutine to handle input from the channel
    replWg.Add(1)
    go func() {
        defer replWg.Done()
        for input := range inputChan {
            processLine(input, out, env, latestOutput, mutex, doneChan)
        }
    }()

    replWg.Wait()
}

func processLine(line string, out io.Writer, env *object.Environment, latestOutput *string, mutex *sync.Mutex, doneChan chan struct{}) {
    l := lexer.New(line)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
        printParserErrors(out, p.Errors())
        return
    }

    evaluated := evaluator.Eval(program, env)
    if evaluated != nil {
        output := evaluated.Inspect()
        io.WriteString(out, output)
        io.WriteString(out, "\n")

        // Store the latest output
        mutex.Lock()
        *latestOutput = output
        mutex.Unlock()
    }

    // Signal that the processing is done
    doneChan <- struct{}{}
}

func printParserErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t"+msg+"\n")
    }
}
