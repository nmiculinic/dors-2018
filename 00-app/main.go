package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	folder := os.Getenv("DORS_FOLDER")
	if folder == "" {
		var err error
		folder, err = os.Getwd()
		if err != nil {
			log.Panicln(err)
		}
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Panicln(err)
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Got request from %v for %v", request.RemoteAddr, request.RequestURI)
		fmt.Fprintln(writer, "hello world from", hostname)
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		fname := folder + "/healthz"
		_, err := os.Open(fname)
		if os.IsNotExist(err) {
			writer.WriteHeader(500)
			fmt.Fprintln(writer, "No file", fname, "found")
		}
	})

	http.HandleFunc("/cnt", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Got request from %v for %v", request.RemoteAddr, request.RequestURI)
		fname := folder + "/cnt"
		cnt := 0
		f, err := os.Open(fname)
		switch {
		case os.IsNotExist(err):
			break
		case err == nil:
			if _, err = fmt.Fscanf(f,"%d", &cnt); err != nil {
				log.Panicln(err)
			}
			if err = f.Close(); err != nil {
				log.Panicln(err)
			}
		default:
			log.Panicln(err)
		}

		cnt++
		f, err = os.Create(fname)
		if err != nil {
			log.Panicln(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Panicln(err)
			}
		}()

		if _, err = fmt.Fprintln(io.MultiWriter(f, writer), cnt); err != nil {
			log.Panicln(err)
		}
	})

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Listening to :8080")
	http.ListenAndServe(":8080", nil)
}
