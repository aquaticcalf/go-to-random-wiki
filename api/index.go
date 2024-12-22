package handler

import (
	"net/http"
	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/*", func(context *Context) {
		http.Redirect(context.Writer, context.Request, "https://en.wikipedia.org/wiki/Special:Random", http.StatusFound)
	})

	server.Handle(w, r)
}

