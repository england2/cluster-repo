package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	targetURL := os.Getenv("SERVICE1_URL")
	if targetURL == "" {
		targetURL = "http://container1:4872/"
	}

	client := &http.Client{Timeout: 5 * time.Second}
	values := []string{"1", "2", "3", "4", "5", "foo"}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		output := values[rng.Intn(len(values))]

		resp, err := client.Post(targetURL, "text/plain", bytes.NewBufferString(output))
		if err != nil {
			log.Printf("failed to send %q to %s: %v", output, targetURL, err)
		} else {
			resp.Body.Close()
			log.Printf("sent %q to %s (status=%s)", output, targetURL, resp.Status)
		}

		time.Sleep(1 * time.Second)
	}
}
