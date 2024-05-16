package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func ShowId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["id"])
	u := r.URL.Query()
	fmt.Println(u)
	extractedUrl := u.Get("url")
	fmt.Println(extractedUrl)
	http.Redirect(w, r, extractedUrl, http.StatusMovedPermanently)
	fmt.Println("After redirection")

}
