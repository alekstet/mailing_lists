package main

import (
	"log"

	"github.com/alekstet/mailing_lists/sender/cmd"
	"github.com/alekstet/mailing_lists/sender/conf"
)

func main() {
	cnf, err := conf.Cnf()
	if err != nil {
		log.Fatalf("error with config: %s", err)
	}
	cmd.Run(cnf)
}
