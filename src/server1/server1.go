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
	http.HandleFunc("/", handler) //each request call handler
	http.HandleFunc("/count", counter, *http.Request)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echos the Path component of the reqeusted URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
