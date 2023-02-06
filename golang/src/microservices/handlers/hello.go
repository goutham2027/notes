package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// log.Println("hello world")
	// time.Sleep(1 * time.Second)
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		// return
		// we can use http.Error
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
		// log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	}
}
