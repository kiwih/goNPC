package npcserver

import (
	"strings"

	"github.com/gocraft/web"
)

//URL handles the urls in the application
type URL string

const (
	//IndexURL is used for the index/home page
	IndexURL URL = "/"

	//ViewURL is used when viewing a generated NPC
	ViewURL URL = "/view"

	//ViewSavedURL is used when viewing a saved NPC
	ViewSavedURL URL = "/view/:npcId"

	//SaveURL is used when saving the currently generated NPC
	SaveURL URL = "/save"
)

//String returns the URL as a string (to satisfy the Stringer interface)
func (u URL) String() string {
	return string(u)
}

//Make creates a URL
func (u URL) Make(param ...string) string {
	if len(param)%2 != 0 {
		panic("Make URL " + u.String() + " had non-even number of params")
	}

	retStr := u.String()

	for i := 0; i < len(param); i += 2 {
		retStr = strings.Replace(retStr, ":"+param[i], param[i+1], 1)
	}
	return retStr
}

func initRouter() *web.Router {
	rootRouter := web.New(Context{})
	rootRouter.Middleware(web.LoggerMiddleware)
	rootRouter.Middleware(web.ShowErrorsMiddleware)
	rootRouter.Middleware(web.StaticMiddleware("./media/public", web.StaticOption{Prefix: "/public"})) // "public" is a directory to serve files from.)
	rootRouter.Middleware((*Context).AssignStorageMiddleware)
	rootRouter.Middleware((*Context).AssignTemplatesAndSessionsMiddleware)
	rootRouter.Middleware((*Context).GetErrorMessagesMiddleware)
	rootRouter.Middleware((*Context).GetNotificationMessagesMiddleware)

	//rootRouter web paths
	rootRouter.Get(IndexURL.String(), (*Context).IndexHandler)

	return rootRouter
}
