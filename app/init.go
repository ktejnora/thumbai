// aah application initialization - configuration, server extensions, middleware's, etc.
// Customize it per application use case.

package main

import (
	"aahframe.work/aah"
	"gorepositree.com/app/data"
	"gorepositree.com/app/proxypass"
	"gorepositree.com/app/security"
	"gorepositree.com/app/vanity"
)

func init() {

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Server Extensions
	// Doc: https://docs.aahframework.org/server-extension.html
	//
	// Best Practice: Define a function with meaningful name in a package and
	// register it here. Extensions function name gets logged in the log,
	// its very helpful to have meaningful log information.
	//
	// Such as:
	//    - Dedicated package for config loading
	//    - Dedicated package for datasource connections
	//    - etc
	//__________________________________________________________________________

	// Event: OnInit
	// Doc: https://docs.aahframework.org/server-extension.html#event-oninit
	//
	// aah.OnInit(config.LoadRemote)

	// Event: OnStart
	// Doc: https://docs.aahframework.org/server-extension.html#event-onstart
	//
	aah.OnStart(SubscribeHTTPEvents)
	aah.OnStart(SubscribeWebSocketEvents)
	aah.OnStart(data.Load, 1)
	aah.OnStart(vanity.Load, 2)
	aah.OnStart(proxypass.Load, 3)

	// Event: OnPreShutdown
	// Doc: https://docs.aahframework.org/server-extension.html#event-onpreshutdown
	//
	// aah.OnPreShutdown(notify.AnnounceImGonnaShutdown)

	// Event: OnPostShutdown
	// Doc: https://docs.aahframework.org/server-extension.html#event-onpostshutdown
	//
	aah.OnPostShutdown(data.Persist)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Middleware's
	// Doc: https://docs.aahframework.org/middleware.html
	//
	// Executed in the order they are defined. It is recommended; NOT to change
	// the order of pre-defined aah framework middleware's.
	//__________________________________________________________________________
	aah.AppHTTPEngine().Middlewares(
		aah.RouteMiddleware,
		// aah.CORSMiddleware,
		aah.BindMiddleware,
		// aah.AntiCSRFMiddleware,
		// aah.AuthcAuthzMiddleware,

		//
		// NOTE: Register your Custom middleware's right here
		//

		aah.ActionMiddleware,
	)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Application Error Handler
	// Doc: https://docs.aahframework.org/error-handling.html
	//__________________________________________________________________________
	// aah.SetErrorHandler(AppErrorHandler)

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Custom Template Functions
	// Doc: https://docs.aahframework.org/template-funcs.html
	//__________________________________________________________________________
	// aah.AddTemplateFunc(template.FuncMap{
	// // ...
	// })

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Custom Session Store
	// Doc: https://docs.aahframework.org/session.html
	//__________________________________________________________________________
	// aah.AddSessionStore("redis", &RedisSessionStore{})

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Custom Value Parser
	// Doc: https://docs.aahframework.org/request-parameters-auto-bind.html
	//__________________________________________________________________________
	// if err := aah.AddValueParser(reflect.TypeOf(CustomType{}), customParser); err != nil {
	//   log.Error(err)
	// }

	//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
	// Add Custom Validation Functions
	// Doc: https://godoc.org/gopkg.in/go-playground/validator.v9
	//__________________________________________________________________________
	// Obtain aah validator instance, then add yours
	// validator := aah.Validator()
	//
	// // Add your validation funcs

}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// HTTP Events
//
// Subscribing HTTP events on app start.
//__________________________________________________________________________

func SubscribeHTTPEvents(_ *aah.Event) {
	he := aah.AppHTTPEngine()

	// Event: OnRequest
	// Doc: https://docs.aahframework.org/server-extension.html#event-onrequest
	//
	// he.OnRequest(repo.Switch)

	// Event: OnPreReply
	// Doc: https://docs.aahframework.org/server-extension.html#event-onprereply
	//
	// he.OnPreReply(myserverext.OnPreReply)

	// Event: OnHeaderReply
	// Doc: https://docs.aahframework.org/server-extension.html#event-onheaderreply
	//
	// he.OnHeaderReply(myserverext.OnHeaderReply)

	// Event: OnPostReply
	// Doc: https://docs.aahframework.org/server-extension.html#event-onpostreply
	//
	// he.OnPostReply(myserverext.OnPostReply)

	// Event: OnPreAuth
	// Doc: https://docs.aahframework.org/server-extension.html#event-onpreauth
	//
	// he.OnPreAuth(myserverext.OnPreAuth)

	// Event: OnPostAuth
	// Doc: https://docs.aahframework.org/server-extension.html#event-onpostauth
	//
	he.OnPostAuth(security.PostAuthEvent)
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// WebSocket Events
// Doc: https://docs.aahframework.org/websocket.html#events
//
// Subscribing WebSocket events on app start.
//__________________________________________________________________________

func SubscribeWebSocketEvents(_ *aah.Event) {
	// wse := aah.AppWSEngine()

	// Custom ID Generator
	//
	// wse.SetIDGenerator(websockets.MyCustomIDGenerator)

	// Event: OnPreConnect
	//
	// wse.OnPreConnect(mywebsockets.HandleEvents)

	// Event: OnPostConnect
	//
	// wse.OnPostConnect(mywebsockets.HandleEvents)

	// Event: OnPostDisconnect
	//
	// wse.OnPostDisconnect(mywebsockets.HandleEvents)

	// Event: OnError
	//
	// wse.OnError(mywebsockets.HandleEvents)
}