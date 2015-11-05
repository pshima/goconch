package main

import (
	"log"
	"math"
	"net"
	"os"
	"time"
)

const (
	interval = 5 // How often in seconds we wait until looping
)

// History to keep track of statistics over time
// Still in early phases for statistics so Stats isn't used yet
type History struct {
	Stats map[string]map[string]float64
	Tests []TestResult
}

// TestResult to track our individual results
type TestResult struct {
	Name     string
	Result   float64
	Duration time.Duration
}

// runCheck does the actual tcp connection
func (e *Endpoint) runCheck() (string, float64, time.Duration) {
	status := 0.0
	starttime := time.Now()
	realTimeout := time.Duration(e.Timeout) * time.Second
	conn, err := net.DialTimeout(e.CheckType, e.Address, realTimeout)
	elapsedtime := time.Since(starttime)
	if err != nil {
		log.Printf("[FAIL] %s failed after %s with error: %s", e.Address, elapsedtime, err)
	} else {
		log.Printf("[SUCCESS] %s: %s", e.Address, elapsedtime)
		conn.Close()
		status = 1.0
	}
	return e.Address, status, elapsedtime
}

// Runner is used to process the input queue and send to the output queue
func Runner(input <-chan Endpoint, output chan<- TestResult) {
	for r := range input {
		endpoint, status, elapsedtime := r.runCheck()
		output <- TestResult{Name: endpoint, Result: status, Duration: elapsedtime}
	}
}

// Calculate an average from a slice
func avg(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return 0.0
	}
	var final float64
	for _, x := range n {
		if final == 0 {
			final = x
		} else {
			final *= x
		}
	}
	return math.Pow(final, 1/float64(l))
}

// StatsRunner processes the output queue and computes statistics
func StatsRunner(output <-chan TestResult, h *History) {
	for r := range output {
		h.Tests = append(h.Tests, r)
		availability := []float64{}
		for _, s := range h.Tests {
			if s.Name == r.Name {
				availability = append(availability, s.Result)
			}
		}
		log.Println("[INFO]", r.Name, "Availability:", avg(availability)*100)
	}
}

func main() {
	// Setup our endpoints from the json file argument
	endpoints := new(Endpoints)
	err := endpoints.FromJSONFile(os.Args)
	if err != nil {
		log.Println("Error loading file from JSON, %v", err)
		return
	}

	// Make a channel for input work to do with the length of
	// 3 times the number of things to check
	input := make(chan Endpoint, len(endpoints.Checklist)*3)
	output := make(chan TestResult, len(endpoints.Checklist)*3)

	// Create a statistics tracker to hold our data
	h := new(History)

	// Launch our workers that read from the queue
	for i := 0; i < len(endpoints.Checklist); i++ {
		go Runner(input, output)
	}

	// Launch a single stats runner
	go StatsRunner(output, h)

	log.Printf("[INFO] Waiting %v seconds before starting", interval)
	log.Printf("[INFO] There are %v total endpoints to be checked", len(endpoints.Checklist))

	// Tick at every X seconds
	ticker := time.NewTicker(time.Second * interval)
	for range ticker.C {
		log.Println("[INFO] Input Queue Length Is:", len(input))
		log.Println("[INFO] Output Queue Length Is:", len(output))
		// Cheap hack to only keep trailing 60 seconds
		if len(h.Tests) > len(endpoints.Checklist)*20 {
			h.Tests = []TestResult{}
		}
		// Put some jobs in the queue
		for _, e := range endpoints.Checklist {
			input <- e
		}

	}
	defer ticker.Stop()
}
