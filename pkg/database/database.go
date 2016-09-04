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
	rhost := os.Getenv(RETHINKDB_HOST)
	if rhost == "" {
		rhost = "127.0.0.1"
	}
	log.Info("Rethink_Host: ", os.Getenv(RETHINKDB_HOST))

	rport := os.Getenv(RETHINKDB_PORT)
	if rport == "" {
		rport = "28015"
	}
	log.Info("Rethink_Host: ", os.Getenv(RETHINKDB_PORT))


	url := rhost+":"+rport
	log.Info("Rethinkdb address: ",url)

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
	log.Info("Databse was created!")

	// Create table: customer
	if _, err := r.DB(dbName).TableCreate("customer").Run(Session); err != nil {
		log.Error(err)
	}
	log.Info("Table was created!")

	// Create table: cluster
	if _, err := r.DB(dbName).TableCreate("cluster").Run(Session); err != nil {
		log.Error(err)
	}
}