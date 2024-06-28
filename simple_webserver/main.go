package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func greetPage(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!",name)
}


func notFoundPage(w http.ResponseWriter,r *http.Request){
	http.NotFound(w,r)
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/greet", greetPage)
}

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	// Wrap the default mux to handle 404 errors
	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
		if r.URL.Path != "/" && r.URL.Path != "/greet" {
			notFoundPage(w, r)
		}
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", wrappedMux)
}