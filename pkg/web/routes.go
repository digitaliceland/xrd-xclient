package web

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, app.session.Enable)
	//dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()
	mux.Get("/", standardMiddleware.ThenFunc(app.timeForm))
/*	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))
*/
	mux.Get("/time", standardMiddleware.ThenFunc(app.timeForm))
	mux.Post("/time", standardMiddleware.ThenFunc(app.timePost))
	mux.Get("/ping", standardMiddleware.ThenFunc(app.pingForm))
	mux.Post("/ping", standardMiddleware.ThenFunc(app.pingPost))
	mux.Get("/continuous", standardMiddleware.ThenFunc(app.continuous))
	mux.Post("/timeservice", standardMiddleware.ThenFunc(app.timeService))

	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
