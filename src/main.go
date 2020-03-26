package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello!! Ningchan")

	http.HandleFunc("/hello", hello)
	fmt.Printf("Starting server port :8080 for testing HTTP GET...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/* hello */
type Hello struct {
	Hello string `json:"hello"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Server", "Ning Mock Server (GoLang): Hello")
		w.WriteHeader(200)
		hello := Hello{
			"Ningchan",
		}
		js, err := json.Marshal(hello)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "POST":
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
