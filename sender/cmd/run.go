package cmd

import (
	"net"
	"net/http"

	"github.com/alekstet/mailing_lists/sender/api/routes"
	"github.com/alekstet/mailing_lists/sender/conf"
)

func Run(cnf *conf.Config) {
	store := conf.New()

	err := store.InitDB(cnf)
	if err != nil {
		store.Log.Fatal(err)
	}

	listener, err := net.Listen("tcp", cnf.Ip+":"+cnf.Port)
	if err != nil {
		store.Log.Fatal(err)
	}

	routes.Routes(*store)

	server := &http.Server{
		Addr:    cnf.Ip + cnf.Port,
		Handler: store.Routes,
	}
	store.Log.Info("Server is running on ", cnf.Ip+":"+cnf.Port)
	store.Log.Fatal(server.Serve(listener))
}
