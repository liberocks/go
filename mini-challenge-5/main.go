package main

import (
	"net/http"

	"github.com/liberocks/go/mini-challenge-5/handler"
)

func main() {

	http.HandleFunc("/", handler.GetSignInPage)
	http.HandleFunc("/profile", handler.GetProfilePage)

	http.HandleFunc("/api/sign-in", handler.PostSignIn)
	http.HandleFunc("/api/sign-out", handler.PostSignOut)

	http.ListenAndServe(":8080", nil)
}
