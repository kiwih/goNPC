package npcserver

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore

	templates *template.Template
	decoder   = schema.NewDecoder() //this initializes the schema (HTML form decoding) engine
)

//StartServer is for starting the npc generation server
func StartServer(templatesDir string, publicDir string, serverAddress string, cookieStoreSalt string) {

	templates = template.Must(template.New("").Funcs(nil).ParseGlob(templatesDir + "/*")) //this initializes the template engine

	store = sessions.NewCookieStore([]byte(cookieStoreSalt))

	//gob is used when we save failed form structs to the session

	router := initRouter(publicDir)

	log.Println("Server running at " + serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Println("Error:", err.Error())
	}
}
