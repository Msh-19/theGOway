package handlers

import(
	"fmt"
	"net/http"
	"url_shortner/storage"
	"url_shortner/utils"
)

func ShortenURL(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return 
	}

	longURL := r.FormValue("url")
	if longURL == ""{
		http.Error(w, "URL is required", http.StatusBadRequest)
		return 
	}

	ShortURL := utils.GenerateShortURL()
	storage.SaveURLMapping(ShortURL,longURL)

	fmt.Fprintf(w, "Short URL: http://localhost:8080/r/%s\n", ShortURL)
}