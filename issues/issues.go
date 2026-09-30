package issues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const baseURL = "https://api.github.com/search/issues"
const userURL = "https://api.github.com/users/"

var githubToken = os.Getenv("GITHUB_TOKEN")

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Body      string
	User      *User
	Milestone *Milestone
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	PublicRepos int    `json:"public_repos,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Type        string `json:"type,omitempty"`
	Location    string `json:"location,omitempty"`
	Email       string `json:"email,omitempty"`
	BIO         string `json:"bio,omitempty"`
	HTMLURL     string `json:"html_url"`
}

type Milestone struct {
	Number      int
	Title       string
	Description string
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, q, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get issues, error with code: %d", resp.StatusCode)
	}

	issues := IssueSearchResult{}

	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		return nil, err
	}

	return &issues, nil
}
