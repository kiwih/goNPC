package npcserver

import "github.com/gocraft/web"

//SetErrorMessage allows for a handler to set an error message as a "Flash" message which can be shown to the user in a later request
//(via a different handler) - it stores them in a session variable
func (c *Context) SetErrorMessage(rw web.ResponseWriter, req *web.Request, err string) {
	session, _ := c.Store.Get(req.Request, "error-messages")
	session.AddFlash(err)
	session.Save(req.Request, rw)
}

//SetNotificationMessage allows for a handler to set a notification message as a "Flash" message which can be shown to the user in a later request
//(via a different handler) - it stores them in a session variable
func (c *Context) SetNotificationMessage(rw web.ResponseWriter, req *web.Request, notification string) {
	session, _ := c.Store.Get(req.Request, "notification-messages")
	session.AddFlash(notification)
	session.Save(req.Request, rw)
}

//SetFailedRequestObject allows us to store a bad request from a form (eg not meeting the regex for the NHI parameter of system.Patient)
//so it can be recalled later for them to amend it
func (c *Context) SetFailedRequestObject(rw web.ResponseWriter, req *web.Request, requestedObject interface{}) {
	session, _ := c.Store.Get(req.Request, "error-form-requests")
	session.AddFlash(requestedObject)
	session.Save(req.Request, rw)
}

//CheckFailedRequestObject returns just one "flash" failed request object for a session. Any other request objects that were stored will be
//removed without retrieval. It allows for users to amend bad forms without needing to retype all the data
func (c *Context) CheckFailedRequestObject(rw web.ResponseWriter, req *web.Request) interface{} {
	session, _ := c.Store.Get(req.Request, "error-form-requests")
	flashes := session.Flashes()
	session.Save(req.Request, rw)

	if len(flashes) > 0 {
		//again, note that only the first one is returned. All other forms will be discarded.
		return flashes[0]
	}
	return nil
}
