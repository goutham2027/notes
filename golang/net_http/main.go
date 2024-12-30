package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

var cacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", handleUser)

	fmt.Println("Server listening to 8081")
	http.ListenAndServe(":8081", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)

}

func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// r.PathValue supports > go 1.22
	id := r.PathValue("id")
	// id, err := getUrlParams(r, "id")
	// if err != nil {
	// 	http.Error(w, "id param is missing", http.StatusBadRequest)
	// 	return
	// }

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// lock reading - to make the operation threadsafe
	cacheMutex.RLock()
	user, ok := userCache[idInt]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	cacheMutex.RUnlock()

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// id, err := getUrlParams(r, "id")
	// if err != nil {
	// 	http.Error(w, "id param is missing", http.StatusBadRequest)
	// 	return
	// }

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// lock reading - to make the operation threadsafe
	cacheMutex.Lock()
	_, ok := userCache[idInt]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	delete(userCache, idInt)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

// func getUrlParams(r *http.Request, param string) (string, error) {
// 	parts := strings.Split(r.URL.Path, "/")
// 	if len(parts) > 2 && parts[1] == param {
// 		return parts[2], nil
// 	} else {
// 		return "", errors.New("param not found")
// 	}
// }
