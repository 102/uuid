package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const sourceHeaderName string = "X-Uuid-Request-Source"
const defaultSourceValue string = "unknown"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	source := defaultSourceValue
	header := r.Header[sourceHeaderName]
	if len(header) > 0 {
		source = header[0]
	}
	res := uuid.New().String()
	log.Printf("uuid generated: %s; source: %s\n", res, source)
	fmt.Fprintf(w, "%s", res)
}
