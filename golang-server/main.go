package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func loginFunction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "login.html") // Đảm bảo file này tồn tại cùng thư mục chạy code
	case "POST":
		if r.FormValue("username") == "admin" && r.FormValue("password") == "admin" {
			fmt.Fprintf(w, "Login successful!")	
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloFunction)
	http.HandleFunc("/login", loginFunction)
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Đăng ký route này SAU cùng

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
