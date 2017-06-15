package npcserver

import "github.com/gocraft/web"

//MIDDLEWARE

//AssignStorageMiddleware creates the link to the storage
func (c *Context) AssignStorageMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Storage = nil
	next(rw, req)
}

//AssignTemplatesAndSessionsMiddleware creates the context's link to the templates and sessions handlers
func (c *Context) AssignTemplatesAndSessionsMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Store = store
	next(rw, req)
}

//GetErrorMessagesMiddleware returns any flash error messages that have been saved. Upon retrieving them, they will be deleted from the session
//(as they are "flash" session variables)
func (c *Context) GetErrorMessagesMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	session, _ := c.Store.Get(req.Request, "error-messages")
	flashes := session.Flashes()
	session.Save(req.Request, rw)

	if len(flashes) > 0 {
		//it is not possible in go to cast from []interface to []string
		strings := make([]string, len(flashes))
		for i := range flashes {
			strings[i] = flashes[i].(string)
		}
		c.ErrorMessages = strings
	}
	next(rw, req)
}

//GetNotificationMessagesMiddleware returns any flash notification messages that have been saved. Upon retrieving them, they will be deleted from the session
//(as they are "flash" session variables)
func (c *Context) GetNotificationMessagesMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	session, _ := c.Store.Get(req.Request, "notification-messages")
	flashes := session.Flashes()
	session.Save(req.Request, rw)

	if len(flashes) > 0 {
		//it is not possible in go to cast from []interface to []string
		strings := make([]string, len(flashes))
		for i := range flashes {
			strings[i] = flashes[i].(string)
		}
		c.NotificationMessages = strings
	}
	next(rw, req)
}
