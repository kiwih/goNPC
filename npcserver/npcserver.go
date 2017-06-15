package npcserver

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore

	templates = template.Must(template.New("").Funcs(nil).ParseGlob(getExecLoc() + "/media/templates/*")) //this initializes the template engine
	decoder   = schema.NewDecoder()                                                                       //this initializes the schema (HTML form decoding) engine
)

func getExecLoc() string {
	loc, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path.Dir(loc)
}

//StartServer is for starting the npc generation server
func StartServer(serverAddress string, cookieStoreSalt string) {

	store = sessions.NewCookieStore([]byte(cookieStoreSalt))

	//gob is used when we save failed form structs to the session

	router := initRouter()

	log.Println("Server running at " + serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Println("Error:", err.Error())
	}
}
