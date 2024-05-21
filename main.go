package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "sync"
    "while/repl"
)

type PostData struct {
    Input string `json:"input"`
}

type ResponseData struct {
    Output string `json:"output"`
}

var (
    latestOutput string
    mutex        sync.Mutex
)

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    })
}

func postHandler(inputChan chan<- string, doneChan chan struct{}) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var data PostData
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        inputChan <- data.Input

        // Wait for the processing to complete
        <-doneChan

        // Ensure latestOutput is updated after processing
        mutex.Lock()
        response := ResponseData{Output: latestOutput}
        mutex.Unlock()

        jsonResponse, err := json.Marshal(response)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
    }
}

func startAPIServer(wg *sync.WaitGroup, inputChan chan<- string, doneChan chan struct{}) {
    defer wg.Done()

    http.Handle("/evaluate", enableCORS(postHandler(inputChan, doneChan)))
    fmt.Println("Starting API server on :6969")
    http.ListenAndServe(":6969", nil)
}

func main() {
    var wg sync.WaitGroup
    inputChan := make(chan string)
    doneChan := make(chan struct{})

    wg.Add(1)
    go startAPIServer(&wg, inputChan, doneChan)

    wg.Add(1)
    go repl.Start(os.Stdin, os.Stdout, &wg, inputChan, &latestOutput, &mutex, doneChan)

    wg.Wait()
}
