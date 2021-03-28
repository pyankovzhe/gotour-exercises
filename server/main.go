package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	PORT = "8000"
	HOST = "localhost"
)

var mu sync.Mutex
var count int

func init() {
	log.Printf("Server is running on %s", fmt.Sprintf("%s:%s", HOST, PORT))
}

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

	users := []user{
		{Name: "User", Email: "user@test.com"},
		{Name: "Admin", Email: "admin@test.com"},
	}

	json.NewEncoder(w).Encode(users)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

func withCounter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		h(w, r)
	}
}

func withLogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
		h(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", withLogRequest(withCounter(handler)))
	mux.HandleFunc("/users", withLogRequest(withCounter(getUsers)))
	mux.HandleFunc("/metrics", metrics)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), mux))
}
