package web

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, app.session.Enable)

	mux := pat.New()
	mux.Get("/", standardMiddleware.ThenFunc(app.timeForm))
	mux.Get("/time", standardMiddleware.ThenFunc(app.timeForm))
	mux.Post("/time", standardMiddleware.ThenFunc(app.timePost))
	mux.Get("/ping", standardMiddleware.ThenFunc(app.pingForm))
	mux.Post("/ping", standardMiddleware.ThenFunc(app.pingPost))
	mux.Get("/continuous", standardMiddleware.ThenFunc(app.continuous))
	mux.Post("/timeservice", standardMiddleware.ThenFunc(app.timeService))

	return standardMiddleware.Then(mux)
}
