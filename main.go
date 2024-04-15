package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error, %v", err)
	}
	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404, not found\n", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported\n", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	port := "8080"
	fmt.Printf("starting server at %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
