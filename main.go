package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if len(queryParams) == 0 {
		fmt.Fprintf(w, "No query parameters provided\n")
		return
	}

	fmt.Fprintf(w, "Echo - Query Parameters:\n")
	fmt.Fprintf(w, "========================\n\n")

	for key, values := range queryParams {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	}
}

func main() {
	http.HandleFunc("/", echoHandler)

	port := "8080"
	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("Try: http://localhost:%s/?name=John&age=30", port)

	addr := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
