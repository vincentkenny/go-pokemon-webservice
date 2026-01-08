package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// global variables
var (
	totalChecks int
	mu          sync.Mutex
	counterCh   = make(chan counterCmd)
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type counterCmd struct {
	delta int
	reply chan int
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go HTTP Server")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	logWithTime("start whole process")
	ch := make(chan string, 2)

	go func() {
		ch <- slowCheck("db", 500*time.Millisecond)
	}()

	go func() {
		ch <- slowCheck("cache", 500*time.Millisecond)
	}()

	go func() {
		ch <- externalSlowCheck("external-service")
	}()

	db := <-ch
	cache := <-ch
	external := <-ch

	fmt.Println(getTotalChecks())

	response := map[string]string{
		"db":       db,
		"cache":    cache,
		"external": external,
	}

	logWithTime("end whole process")

	json.NewEncoder(w).Encode(response)
}

func getTotalChecks() int {
	reply := make(chan int)
	counterCh <- counterCmd{delta: 0, reply: reply}
	return <-reply
}

func slowCheck(name string, delay time.Duration) string {
	logWithTime("starting slowCheck" + name)
	time.Sleep(delay)

	// //mutex way of preventing race condition
	// mu.Lock()
	// totalChecks++
	// mu.Unlock()

	//channel way of preventing race condition
	counterCh <- counterCmd{delta: 1}

	logWithTime("finished slowCheck" + name)
	return name + " ok"
}

func externalSlowCheck(name string) string {
	logWithTime("starting externalSlowCheck")
	time.Sleep(2000 * time.Millisecond)

	// //mutex way of preventing race condition
	// mu.Lock()
	// totalChecks++
	// mu.Unlock()

	//channel way of preventing race condition
	counterCh <- counterCmd{delta: 1}

	logWithTime("finished externalSlowCheck")
	return name + "ok"
}

func logWithTime(msg string) {

	fmt.Println(time.Now().Format("15:04:05.00"), msg)
}

func startCounter() {
	go func() {
		total := 0
		for cmd := range counterCh {
			total += cmd.delta
			if cmd.reply != nil {
				cmd.reply <- total
			}
		}
	}()
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)

	startCounter()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
