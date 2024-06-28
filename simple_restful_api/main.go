package main

import(
	"fmt"
	"log"
	"net/http"
	"simple_restful_api/handlers"
)

func main()  {
	http.HandleFunc("/books", handlers.GetBooks)
	http.HandleFunc("/books/",handlers.GetBook)
	http.HandleFunc("/books",handlers.CreateBook)
	http.HandleFunc("/books/",handlers.UpdateBook)
	http.HandleFunc("/books",handlers.DeleteBook)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatalf("Error starting server: %s\n", err)
	}
}