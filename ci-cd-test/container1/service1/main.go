package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}

		output := string(body)
		fmt.Printf("received: %s\n", strings.TrimSpace(output))

		if strings.TrimSpace(output) == "foo" {
			fmt.Println("got foo")
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(body); err != nil {
			log.Printf("failed to write response body: %v", err)
		}
	})

	addr := "0.0.0.0:4872"
	log.Printf("service1 listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
