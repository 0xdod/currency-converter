package main

import (
	"log"
	"net/http"
)

func main() {
	server := &Server{}
	log.Println("Server started on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server))
}
