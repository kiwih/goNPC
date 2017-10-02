package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/kiwih/npc-gen/npcserver"
	"github.com/kiwih/npc-gen/questgen"
)

func getExecLoc() string {
	loc, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path.Dir(loc)
}

const (
	publicDir    = "/media/public"
	templatesDir = "/media/templates"
	dataDir      = "/media/data"
)

func main() {
	//Enable logger
	logFileName := os.Getenv("LOG_FILE_NAME")
	if len(logFileName) == 0 {
		logFileName = "npcserver.log"
	}

	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("Can't open log file: " + err.Error())
	}
	log.SetOutput(io.MultiWriter(f, os.Stdout))
	defer f.Close()

	// On platforms like heroky/dokku this should be PORT, not HTTP_PORT. On Azure this should be HTTP_PLATFORM_PORT
	serverAddress := ":" + os.Getenv("HTTP_PORT")
	if serverAddress == ":" {
		log.Println("$HTTP_PORT was not set, defaulting to 3000")
		serverAddress = ":3000"
	}

	cookieStoreSalt := os.Getenv("COOKIE_STORE_SALT")
	if len(cookieStoreSalt) == 0 {
		log.Println("$COOKIE_STORE_SALT was not set, defaulting to 'SUPER_SECRET_SALT'.")
		cookieStoreSalt = "SUPER_SECRET_SALT"
	}

	execLoc := getExecLoc()

	//fmt.Println(npcgen.BanditCaptain.String())

	basicQuest := questgen.GenerateQuest(
		questgen.GetBasicQuestLocations(),
		questgen.GetBasicQuestNPCs(),
		questgen.GetBasicQuestItems(),
		4,
	)

	fmt.Printf("%+v\n", basicQuest)

	npcserver.StartServer(execLoc+templatesDir, execLoc+publicDir, serverAddress, cookieStoreSalt)
}
