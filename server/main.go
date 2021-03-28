package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func metrics(w http.ResponseWriter, r *http.Request) {
	metrics := map[string]int{
		"http_requests_total": count,
	}

	json.NewEncoder(w).Encode(metrics)
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	users := []user{
		{Name: "User", Email: "user@test.com"},
		{Name: "Admin", Email: "admin@test.com"},
	}

	json.NewEncoder(w).Encode(users)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/users", getUsers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
