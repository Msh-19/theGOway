package handlers

import(
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the URL Shortener Service")
}