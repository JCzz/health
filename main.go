package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var initialHealthStatus string
var messageOrginal string
var message string
var toggleTime int
var errorReadingToggleInterval error

func ticking() {
	if toggleTime == -1 {
		message = messageOrginal + ": toggleTime = -1"
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	// defer ticker.Stop()

	countdown := toggleTime
	go func() {
		for range ticker.C {
			// fmt.Println("tick")
			countdown--
			if countdown == 0 {
				message = messageOrginal + ": done"
			} else {
				message = messageOrginal + ": " + strconv.Itoa(countdown)
			}

			if countdown == 0 {
				if initialHealthStatus == "bad" {
					initialHealthStatus = "good"
				} else if initialHealthStatus == "good" {
					initialHealthStatus = "bad"
				}
				ticker.Stop()
				break
			}
		}
	}()
}

func healthStatusHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Health Status: %s", initialHealthStatus)
	fmt.Fprint(w, initialHealthStatus)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Health Status: %s", initialHealthStatus)
	fmt.Fprint(w, message)
}

func main() {

	initialHealthStatus = os.Getenv("INITIAL_HEALTH_STATUS")
	if initialHealthStatus == "" {
		initialHealthStatus = "not set"
	}

	toggleIntervalStr := os.Getenv("TOGGLE_INTERVAL")
	toggleTime, errorReadingToggleInterval = strconv.Atoi(toggleIntervalStr)
	if errorReadingToggleInterval != nil {
		toggleTime = -1
	}
	messageOrginal = os.Getenv("MESSAGE")

	ticking()

	http.HandleFunc("/health", healthStatusHandler)
	http.HandleFunc("/", messageHandler)

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
