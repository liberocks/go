package main

import (
	"fmt"
	"net/http"

	"github.com/liberocks/go/mini-challenge-5/handler"
)

func main() {

	http.HandleFunc("/", handler.GetSignInPage)
	http.HandleFunc("/profile", handler.GetProfilePage)

	http.HandleFunc("/api/sign-in", handler.PostSignIn)
	http.HandleFunc("/api/sign-out", handler.PostSignOut)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
