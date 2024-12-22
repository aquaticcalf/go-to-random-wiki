package handler

import (
    "net/http"
    . "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    server := New()
    server.GET("/*", func(context *Context) {
        http.Redirect(context.Writer, r, "https://en.wikipedia.org/wiki/Special:Random", http.StatusMovedPermanently)
    })
    server.Handle(w, r)
}
