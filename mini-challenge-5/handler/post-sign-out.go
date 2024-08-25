package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/liberocks/go/mini-challenge-5/repository"
)

func PostSignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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

	// Delete session
	err = repository.SessionRepository.Delete(repository.SessionRepository{}, session.Session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie = &http.Cookie{Name: "session", Value: "",  Expires: time.Unix(0, 0), MaxAge: -1, HttpOnly: false, Path: "/"}
	http.SetCookie(w, cookie)

	tpl, err := template.ParseFiles("template/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)

}
