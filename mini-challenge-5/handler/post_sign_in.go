package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/liberocks/go/mini-challenge-5/model"
	"github.com/liberocks/go/mini-challenge-5/repository"
	"github.com/liberocks/go/mini-challenge-5/util"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decoding request payload
	var payload SignInRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := payload.Email
	password := payload.Password

	// Fetch user by email
	user, err := repository.UserRepository.GetByEmail(repository.UserRepository{}, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Fetch credential by user ID
	credential, err := repository.CredentialRepository.GetById(repository.CredentialRepository{}, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Check if password is correct
	if credential.Password != password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create session and put it in cookie
	session := model.Session{
		UserId:  user.Id,
		Session: util.GenerateRandomString(32),
	}
	expire := time.Now().Add(60 * time.Minute)
	cookie := http.Cookie{Name: "session", Value: session.Session, Expires: expire, MaxAge: 86400, HttpOnly: false, Path: "/"}

	repository.SessionRepository.Create(repository.SessionRepository{}, session)
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(user)
}
