package storage

import(
	"errors"
)

var urlMap = make(map[string]string)

func SaveURLMapping(shortURL, longURL string)  {
	urlMap[shortURL] = longURL
}

func GetLongURL(shortURL string) (string, error) {
	longURL, exists := urlMap[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}
	return longURL, nil
}