package npcserver

import (
	"net/http"

	"github.com/gocraft/web"
	"github.com/gorilla/sessions"
)

//AnyStorer is the interface that is made up of all other storage interfaces
type AnyStorer interface {
}

//Context is used in all requests
type Context struct {
	ErrorMessages        []string
	NotificationMessages []string
	Data                 interface{}
	Store                *sessions.CookieStore
	Storage              AnyStorer
}

//HANDLERS

//IndexHandler is for serving the home/index page
func (c *Context) IndexHandler(rw web.ResponseWriter, req *web.Request) {
	err := templates.ExecuteTemplate(rw, "indexPage", c)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
