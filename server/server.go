package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Welcome to my website! yall")
		http.ServeFile(w, r, "../client/index.html")
	})
	http.HandleFunc("/messages", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("You can read messages.")
		http.ServeFile(w, r, "../client/messages.html")
	})
	// fs := http.FileServer(http.Dir("../client"))
	// http.Handle("/client/", http.StripPrefix("/client/", fs))

	http.ListenAndServe(":8080", nil)
}
