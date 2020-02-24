package web

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
	"html/template"
	"github.com/golangcollege/sessions"
)

type application struct {
	debug    bool
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
/*	snippets interface {
		Insert(string, string, string) (int, error)
		Get(int) (*models.Snippet, error)
		Latest() ([]*models.Snippet, error)
	}
*/
	templateCache map[string]*template.Template
/*	users         interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
		ChangePassword(int, string, string) error
	}
*/}


func StartWWW(addr string) {
	debug := flag.Bool("debug", false, "Enable debug mode")
	secret := "test123"
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		debug:         *debug,
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		//snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		//users:         &mysql.UserModel{DB: db},
	}


	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}