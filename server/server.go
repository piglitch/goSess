package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Serve the index.html file at the root
    r := mux.NewRouter()

    r.HandleFunc("/", showHomePage)
    // Serve other HTML files based on the route
    r.HandleFunc("/messages", showMessages)

    // Serve all other static files (JS, CSS, images, etc.)
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    // http.HandleFunc("/client/static/js/main.js", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Println("Serving messages.html")
    //     http.ServeFile(w, r, "../client/static/js/main.js")
    // })
    //fmt.Println(fs)

    http.Handle("/", r)
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

func showHomePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Serving index.html")
    http.ServeFile(w, r, "../client/index.html")
}

func showMessages(w http.ResponseWriter, r *http.Request){
    fmt.Println("Serving messages.html")
    http.ServeFile(w, r, "../client/messages.html")
}

