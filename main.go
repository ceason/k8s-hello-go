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
	defer observeRequest(time.Now())

	// random sleep to simulate service actually doing something
	time.Sleep(time.Duration(rand.Int63n(50)) * time.Millisecond)

	// "business logic" for service
	fmt.Fprint(w, "Hello <name >!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Printf("Starting server at http://localhost%v.\n", ":8184")
	log.Fatal(http.ListenAndServe(":8184", nil))
}
