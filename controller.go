package main

import (
	"log"
	"net/http"
)

var Users []User

type Mux struct {
}

type User struct {
	Port      string
	Password  string
	Limit     string
	Used      string
	Remaining string
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		LoginPage(w)
		return
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		loginPrompt := CheckLogin(username, password)
		if loginPrompt == "Login Success" {
			DeployCookie(w, username, 0)
			http.Redirect(w, r, "/panel", http.StatusFound)
			return
		}
		LoginResponse(w, loginPrompt)
		return
	}
}

func panel(w http.ResponseWriter, r *http.Request) {
	// invalid cookie
	if !(CheckCookie(r)) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	switch r.Method {
	case "GET":
		Users = GetUsers(SSUSERSPATH, SSTRAFFICPATH)
		PanelMainPage(w, Users)
		return
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		DestoryCookie(w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

// route be public
func (pm *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		login(w, r)
		return
	case "/panel":
		panel(w, r)
		return
	case "/logout":
		logout(w, r)
		return
	}
}

func DefineListenAndServe() {
	mux := &Mux{}
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
