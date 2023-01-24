package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//  serveMux

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))
			// return
			// we can use http.Error
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		// log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")

	})
	http.ListenAndServe(":9090", nil)
}
