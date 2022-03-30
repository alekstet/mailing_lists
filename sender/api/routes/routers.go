package routes

import (
	"net/http"

	"github.com/alekstet/mailing_lists/sender/api/routes/process"
	"github.com/alekstet/mailing_lists/sender/conf"

	"github.com/julienschmidt/httprouter"
)

type R struct {
	Router *httprouter.Router
}

func New() *R {
	return &R{
		Router: httprouter.New(),
	}
}

func Routes(s conf.Store) *httprouter.Router {
	process.New(s).Register()
	return s.Routes
}

func (rr *R) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.Router.ServeHTTP(w, r)
}
