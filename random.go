package main

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://en.wikipedia.org/wiki/Special:Random", http.StatusFound)
}

var Handlerfunc = http.HandlerFunc(Handler)