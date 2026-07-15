package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var serverID int

func init() {
	serverID = rand.Intn(1000000000)
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Hello World from server %d!", serverID)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func main() {
	fmt.Printf("Server %d started\n", serverID)
	err := http.ListenAndServe(":8080", &handler{})
	if err != nil {
		panic(err)
	}
}
