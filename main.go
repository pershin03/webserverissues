package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webservermilestones/issues"
)

type server struct {
	query  string
	result *issues.IssueSearchResult
}

func main() {
	srv := &server{}
	http.HandleFunc("/", srv.issuesHandler)
	http.HandleFunc("/user", srv.userHandler)
	http.HandleFunc("/milestone", srv.milestoneHandler)
	http.HandleFunc("/issue", srv.issueHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) issuesHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		fmt.Fprintln(w, "введите ?q=... в адресной строке")
		return
	}

	if s.result == nil || s.query != q {
		res, err := issues.SearchIssues([]string{q})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.result = res
		s.query = q
	}

	for _, issue := range s.result.Items {
		fmt.Fprintf(w, "#%d %s [%s]\n", issue.Number, issue.Title, issue.State)
	}

}

func (s *server) userHandler(w http.ResponseWriter, r *http.Request) {
	if s.result == nil {
		http.Error(w, "need to download issues on main page", http.StatusBadRequest)
		return
	}

	
}

func (s *server) milestoneHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *server) issueHandler(w http.ResponseWriter, r *http.Request) {
	if s.result == nil {
		http.Error(w, "need to download issues on main page", http.StatusBadRequest)
		return
	}

	q := r.URL.Query().Get("number")
	num, err := strconv.Atoi(q)
	if err != nil {
		http.Error(w, "query parameter is not number", http.StatusBadRequest)
		return
	}

	fNumExist := false
	for _, issue := range s.result.Items {
		if issue.Number == num {
			fNumExist = true
			if issue.User != nil {
				fmt.Fprintf(w, "%s %s %s", issue.State, issue.Title, issue.User.Login)
			} else {
				fmt.Fprintf(w, "%s %s (unknown user)", issue.State, issue.Title)
			}
			break
		}
	}

	if !fNumExist {
		http.Error(w, fmt.Sprintf("issue with number %d not found\n", num), http.StatusNotFound)
		return
	}
}
