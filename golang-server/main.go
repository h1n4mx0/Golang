package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func loginFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.FormValue("username") == "admin" || r.FormValue("password") == "admin" {
		fmt.Fprintf(w, "Login successful!")
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloFunction)
	http.HandleFunc("login", loginFunction)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}