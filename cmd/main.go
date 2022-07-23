package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// if server was stopped with CTRL + C
func SetupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl + C pressed in Terminal")
		os.Exit(0)
	}()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Form could not be parsed.")
		log.Printf("Error occured: %v", err)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	telegram := r.FormValue("telegarm")

	if name == "" || email == "" || telegram == "" {
		fmt.Fprintf(w, "Form was empty.")
		return
	}

	fmt.Fprintf(w, "Form was successfully posted!\n")
	fmt.Fprintf(w, "Name %s\n", name)
	fmt.Fprintf(w, "Email %s\n", email)
	fmt.Fprintf(w, "Telegram %s\n", telegram)
}

func main() {
	SetupCloseHandler()

	var PORT int = 8080
	fileServer := http.FileServer(http.Dir("./../static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("- Starting server at %d\n", PORT)
	if err := http.ListenAndServe(":"+fmt.Sprint(PORT), nil); err != nil {
		log.Fatal(err)
	}
}
