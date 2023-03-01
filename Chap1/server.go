package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counthandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func handler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/favicon.ico" {
		mu.Lock()
		count++
		fmt.Println(count)
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
}

func counthandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d\n", count)
	mu.Unlock()
}
