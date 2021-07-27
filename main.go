package main

import (
	"context"
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

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/protectrepo", ProtectRepoHandler).Methods("POST")
	http.ListenAndServe(":80", r)
}
