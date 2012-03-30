package hello

import (
	"appengine"
	"appengine/user"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", handlerLogin)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}
