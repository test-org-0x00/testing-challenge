package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	response := fmt.Sprintf(`{
  "message": "Hello, World!",
  "hostname": "%s",
  "version": "1.0.0"
}`, hostname)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)
	
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Printf("Visit http://localhost:%s to see the hello world message\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
