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

		output := strings.TrimSpace(string(body))
		fmt.Printf("received: %s\n", output)

		if output == "foo" {
			fmt.Println("got foo")
		}

		w.WriteHeader(http.StatusOK)
	})

	addr := "0.0.0.0:4872"
	log.Printf("service1 listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
