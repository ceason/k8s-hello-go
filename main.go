package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"math/rand"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// collect prometheus metrics
	rpcRequests.Inc()
	start := time.Now()
	defer func() {
		durationMs := float64(time.Since(start).Nanoseconds()) / 1e6
		rpcDurations.Observe(durationMs)
	}()

	time.Sleep(time.Duration(rand.Int63n(50)) * time.Millisecond)

	// "business logic" for service
	fmt.Fprint(w, "Hello <name>!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Printf("Starting server at http://localhost%v.\n", ":8184")
	log.Fatal(http.ListenAndServe(":8184", nil))
}
