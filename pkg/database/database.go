package database

import (
	log "github.com/Sirupsen/logrus"
	r "gopkg.in/dancannon/gorethink.v2"
	"os"
)

const (
	dbName = "mws"
)

const (
	RETHINKDB_HOST	= "RETHINKDB_HOST"
	RETHINKDB_PORT	= "RETHINKDB_PORT"
)

var Session *r.Session

func InitDb() {
	log.Info("Rethink_Host: ", RETHINKDB_PORT)
	url := "172.18.0.2"+os.Getenv(RETHINKDB_PORT)
	if url == "" {
		log.Fatal("Environment varialbe RETHINKDB_URL is empty!")
		return
	}
	log.Info(url)

	dbSession, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	Session = dbSession

	// Create database: mws
	if _, err := r.DBCreate(dbName).Run(Session); err != nil {
		log.Error(err)
	}

	// Create table: customer
	if _, err := r.DB(dbName).TableCreate("customer").Run(Session); err != nil {
		log.Error(err)
	}

	// Create table: cluster
	if _, err := r.DB(dbName).TableCreate("cluster").Run(Session); err != nil {
		log.Error(err)
	}
}