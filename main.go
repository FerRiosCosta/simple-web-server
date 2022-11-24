package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("app-version: 2\n"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Print("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
