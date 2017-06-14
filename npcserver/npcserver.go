package npcserver

import (
	"html/template"
	"os"
	"path"

	"github.com/gorilla/schema"
)

var (
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

func StartServer(serverAddress string) {

	//gob is used when we save failed form structs to the session

	// router := initRouter()

	// log.Println("Server running at " + serverAddress)
	// if err := http.ListenAndServe(serverAddress, router); err != nil {
	// 	log.Println("Error:", err.Error())
	// }
}
