package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handleRequest).Methods("GET")
	r.HandleFunc("/kubecon", handleRequest).Methods("GET")
	return r

}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Valenica from Intuit!")
}
func main() {
	r := Handlers()

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal(err)
	}
}
