package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve the index.html file at the root
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Serving index.html")
        http.ServeFile(w, r, "../client/index.html")
    })

    // Serve other HTML files based on the route
    http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Serving messages.html")
        http.ServeFile(w, r, "../client/messages.html")
    })

    // Serve all other static files (JS, CSS, images, etc.)
    // fs := http.FileServer(http.Dir("../client/static/js/main.js"))
    // http.Handle("../client/static/js/", http.StripPrefix("../client/static/js/", fs))
    http.HandleFunc("/client/static/js/main.js", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Serving messages.html")
        http.ServeFile(w, r, "../client/static/js/main.js")
    })
    //fmt.Println(fs)
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

