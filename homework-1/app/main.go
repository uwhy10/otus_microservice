package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/", getMain)
	http.HandleFunc("/health", getHealth)

	// starting server
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getMain(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, "Hello world!"); err != nil {
		log.Panic(err)
	}
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var j, _ = json.Marshal(StatusResponse{
		Status: "OK",
	})

	if _, err := w.Write(j); err != nil {
		log.Panic(err)
	}
}
