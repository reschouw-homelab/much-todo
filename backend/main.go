package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "[]")
    })
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
