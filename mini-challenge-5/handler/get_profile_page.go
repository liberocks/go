package handler

import (
	// "fmt"

	"html/template"
	"net/http"

	"github.com/liberocks/go/mini-challenge-5/repository"
)

func GetProfilePage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	s := cookie.Value

	// Get session by session ID
	session, err := repository.SessionRepository.GetBySession(repository.SessionRepository{}, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get user by user ID
	user, err := repository.UserRepository.GetById(repository.UserRepository{}, session.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFiles("template/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, user)

}
