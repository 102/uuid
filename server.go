package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", uuid.New().String())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
