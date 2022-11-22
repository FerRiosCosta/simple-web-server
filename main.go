package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Este es la version 1."))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Print("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
