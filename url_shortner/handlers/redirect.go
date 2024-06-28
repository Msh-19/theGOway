package handlers

import(
	"net/http"
	"url_shortner/storage"
)

func RedirectURL(w http.ResponseWriter, r *http.Request)  {
	ShortURL := r.URL.Path[len("/r/"):]
	longURL, err := storage.GetLongURL(ShortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w,r,longURL,http.StatusFound)
}