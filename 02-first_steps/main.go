package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Panicln(err)
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Got request from %v for %v", request.RemoteAddr, request.RequestURI)
		fmt.Fprintln(writer, "hello world from", hostname)
	})
	log.Println("Listening to :8080")
	http.ListenAndServe(":8080", nil)
}
