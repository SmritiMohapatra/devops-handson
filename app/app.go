package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    mt.Fprintf(w, "Hello DevOps")
    fmt.Fprintf(w, "Step 1")
    fmt.Fprintf(w, "Step oh")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
