package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	res := uuid.New().String()
	log.Printf("uuid generated: %s; method: %s\n", res, r.Method)
	fmt.Fprintf(w, "%s", res)
}
