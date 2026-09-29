package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", issuesHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/milestone", milestoneHandler)
	http.HandleFunc("/issue", issueHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func issuesHandler(w http.ResponseWriter, r *http.Request) {

}

func userHandler(w http.ResponseWriter, r *http.Request) {

}

func milestoneHandler(w http.ResponseWriter, r *http.Request) {

}

func issueHandler(w http.ResponseWriter, r *http.Request) {

}
