package process

import (
	"github.com/alekstet/mailing_lists/sender/conf"

	_ "github.com/mattn/go-sqlite3"
)

type S conf.Store

func New(s conf.Store) *S {
	return &S{
		Db:     s.Db,
		Log:    s.Log,
		Routes: s.Routes,
	}
}

func (s *S) Register() {
	s.Routes.POST("/mail", s.AddMail)
	s.Routes.POST("/client", s.AddClient)
	s.Routes.DELETE("/client", s.DeleteClient)
	s.Routes.PUT("/client", s.UpdateClient)
}
