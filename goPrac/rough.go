package main

import (
	"fmt"
    //"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// Struct to define the expected JSON structure
// type User struct {
//     Name  string `json:"name"`
//     Email string `json:"email"`
// }

func main() {
    // Serve the index.html file at the root
    r := mux.NewRouter()

    r.HandleFunc("/", showHomePage)
    // r.HandleFunc("/api/json", getJSON).Methods("GET")
    // r.HandleFunc("/api/json", postJSON).Methods("POST")
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

    // http.Handle("/", r)
    // http.Handle("/messages", r)
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}

func showHomePage(w http.ResponseWriter, r *http.Request){
    // Print the fetched data to the console
    fmt.Println("Serving index.html")
    http.ServeFile(w, r, "../client/index.html")
    
}

// func getJSON(w http.ResponseWriter, r *http.Request){
//     response := map[string]string{
//         "message": "Hello, World!",
//         "status":  "success",
//     }

//     // Set the Content-Type header to application/json
//     w.Header().Set("Content-Type", "application/json")

//     // Encode the response as JSON and send it
//     json.NewEncoder(w).Encode(response)
    
// }

// func postJSON(w http.ResponseWriter, r *http.Request){
//     var user User

//     // For demonstration, let's print the user data to the console
//     fmt.Printf("Received User: %+v\n", user)

//     // Respond with a success message and the user data in JSON format
//     w.Header().Set("Content-Type", "application/json")
//     // Set the Content-Type header to application/json
//     json.NewEncoder(w).Encode(map[string]string{
//         "message": "User created successfully",
//         "name":    user.Name,
//         "email":   user.Email,
//     })
    
// }

func showMessages(w http.ResponseWriter, r *http.Request){
    fmt.Println("Serving messages.html")
    http.ServeFile(w, r, "../client/messages.html")
}

