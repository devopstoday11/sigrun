package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 9000.....")
	log.Fatal(http.ListenAndServe(":9000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("Hello world1"))
		if err != nil {
			log.Println(err)
			return
		}
	})))
}
