package main

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var (
	loginTPL, _ = template.ParseFiles("login.gtpl")
	panelTPL, _ = template.ParseFiles("panel.gtpl")
	salt        = createSalt()
)

// template execute method
func LoginPage(w http.ResponseWriter) {
	loginTPL.Execute(w, "")
}

func LoginResponse(w http.ResponseWriter, loginPrompt string) {
	loginTPL.Execute(w, loginPrompt)
}

func PanelMainPage(w http.ResponseWriter, Users []User) {
	panelTPL.Execute(w, Users)
}

// cookie method
func createSalt() []byte {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	salt := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 8; i++ {
		salt = append(salt, bytes[r.Intn(len(bytes))])
	}
	return salt
}

func md5SaltHash(username string) string {
	key := string(salt) + username
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

func createCookie(username string, maxAge int) (*http.Cookie, *http.Cookie) {
	hash := md5SaltHash(username)
	usernameCookie := http.Cookie{Name: "username", Value: username, MaxAge: maxAge, HttpOnly: true}
	hashCookie := http.Cookie{Name: "hash", Value: hash, MaxAge: maxAge, HttpOnly: true}
	return &usernameCookie, &hashCookie
}

func DeployCookie(w http.ResponseWriter, username string, maxAge int) {
	usernameCookie, hashCookie := createCookie(username, maxAge)
	w.Header().Set("Set-Cookie", usernameCookie.String())
	w.Header().Add("Set-Cookie", hashCookie.String())
}

func DestoryCookie(w http.ResponseWriter) {
	usernameCookie, hashCookie := createCookie("", -1)
	w.Header().Set("Set-Cookie", usernameCookie.String())
	w.Header().Add("Set-Cookie", hashCookie.String())
}

// func GetUsernameFromCookie(r *http.Request) string {
// 	usernameCookie, _ := r.Cookie("username")
// 	return usernameCookie.Value
// }

func CheckCookie(r *http.Request) bool {
	usernameCookie, err1 := r.Cookie("username")
	hashCookie, err2 := r.Cookie("hash")
	if (err1 != nil) || (err2 != nil) {
		return false
	}
	username := usernameCookie.Value
	hash := hashCookie.Value
	if md5SaltHash(username) == hash {
		return true
	}
	return false
}
