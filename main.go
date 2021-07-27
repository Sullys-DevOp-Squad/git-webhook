package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/go-github/v37/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

//define global github client scoped to main package
var client *github.Client

func init() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_KEY")},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client = github.NewClient(tc)
}

func ProtectRepoHandler(w http.ResponseWriter, r *http.Request) {

	req := github.WebHookPayload{}
	err := json.NewDecoder(r.Body).Decode(&req)

	//read json payload
	defer r.Body.Close()

	if err != nil {
		w.Write([]byte("Event did not parse correctly"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check if event is not of type repo created
	if *req.Action != "created" {
		w.Write([]byte("Event did not parse correctly"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//begin creating branch policies with the client
	protectionRequest := &github.ProtectionRequest{
		RequiredStatusChecks: &github.RequiredStatusChecks{
			Strict:   true,
			Contexts: []string{"continuous-integration"},
		},
	}

	_, resp, err := client.Repositories.UpdateBranchProtection(context.Background(), *req.Repo.Owner.Login, *req.Repo.Name, "main", protectionRequest)

	if resp.StatusCode != http.StatusOK {
		w.Write([]byte(err.Error()))
		w.WriteHeader(resp.StatusCode)
		return
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/protectrepo", ProtectRepoHandler).Methods("POST")
	http.ListenAndServe(":80", r)
}
