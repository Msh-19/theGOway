package main

import(
	"fmt"
	"log"
	"net/http"
	"url_shortner/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Homepage)
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/r/", handlers.RedirectURL)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
