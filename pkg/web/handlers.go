package web

import (
	"fmt"
	"github.com/brynjarh/xclient/pkg/forms"
	"net/http"
	"net/url"
)


func (app *application) timeForm(w http.ResponseWriter, r *http.Request) {
	t:="url=http://localhost:80&client=CS/ORG/1111/TestClient&service=CS/ORG/1111/TestService/TEST123"
	v, err := url.ParseQuery(t)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "time.page.gohtml", &templateData{
		Form: forms.New(v),
	})
}

func (app *application) pingForm(w http.ResponseWriter, r *http.Request) {
	t:="url=http://localhost:80&client=CS/ORG/1111/TestClient&service=CS/ORG/1111/TestService/TEST123"
	v, err := url.ParseQuery(t)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "ping.page.gohtml", &templateData{
		Form: forms.New(v),
	})
}

func (app *application) timePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("url", "client", "service")
	form.MatchesPattern("url", forms.UrlRX)
	form.ValidUrl("url")

	if !form.Valid() {
		//for k, v := range form.Errors {
		//	fmt.Printf("key[%s] value[%s]\n", k, v)
		//}
		app.render(w, r, "time.page.gohtml", &templateData{
			Form: form,
			Active: "time",
		})
		return
	}
	u, _ := url.Parse(form.Get("url"))
	// Now u is a valid parsed url for the security server

	// Now we know we have a valid form
	c := &Client{
		BaseURL:  u,
		XRoadClient:  form.Get("client"),
		XRoadService:  form.Get("service"),
		httpClient: http.DefaultClient,
	}
	// We have created the client object and filled in all neccesary data to query the API
	result, req, rep, err := c.do("time")
	// We have called the service and received a reply
	if err != nil {
		result = fmt.Sprintf("ERROR: %s", err)
	}

	app.render(w, r, "time.page.gohtml", &templateData{
		Result: result,
		Form: form,
		Active: "time",
		RequestHeaders: req,
		ReplyHeaders: rep,
	})
}


func (app *application) pingPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("url", "client", "service")
	form.MatchesPattern("url", forms.UrlRX)
	form.ValidUrl("url")

	if !form.Valid() {
		//for k, v := range form.Errors {
		//	fmt.Printf("key[%s] value[%s]\n", k, v)
		//}
		app.render(w, r, "ping.page.gohtml", &templateData{
			Form: form,
			Active: "ping",
		})
		return
	}
	u, _ := url.Parse(form.Get("url"))
	// Now u is a valid parsed url for the security server

	// Now we know we have a valid form
	c := &Client{
		BaseURL:  u,
		XRoadClient:  form.Get("client"),
		XRoadService:  form.Get("service"),
		httpClient: http.DefaultClient,
	}
	// We have created the client object and filled in all neccesary data to query the API
	result, req, rep, err := c.do("ping")
	// We have called the service and received a reply
	if err != nil {
		result = fmt.Sprintf("ERROR: %s", err)
	}

	app.render(w, r, "ping.page.gohtml", &templateData{
		Result: result,
		Form: form,
		Active: "ping",
		RequestHeaders: req,
		ReplyHeaders: rep,
	})
}
