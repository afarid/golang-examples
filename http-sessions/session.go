package main

import (
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	io.WriteString(w, "This is a secret data")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}
