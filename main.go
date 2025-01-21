package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleProcess(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("http://localhost:8081/thisJob", "application/json", nil)
	if err != nil {
		log.Fatal("Failed to send request: ", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response: ", err)
	}

	respString := fmt.Sprintf("Response: %s", respBody)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("HX-Reswap", "afterend")
	w.Write([]byte(respString))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/createJobRun", HandleProcess)

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
