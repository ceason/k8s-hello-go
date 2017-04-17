package main

import (
	"fmt"
	"net/http"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hello Chris and JT. Demo this rocks! You have called me %d times.\n", calls)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8184", nil)
	fmt.Printf("Started server at http://localhost%v.\n", ":8184")
}
