package main

import (
	"log"
	"net/http"
	"os"
	"uuid/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
