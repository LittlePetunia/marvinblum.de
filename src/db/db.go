package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

// Connects to MongoDB database. Will panic on error.
// Pass host and database to use.
func Connect(host, database string) {
	// connect
	log.Print("Connecting to database at ", host)
	session, err := mgo.Dial(host)

	if err != nil {
		panic(err)
	}

	// select db
	log.Print("Selecting database ", database)
	db = session.DB(database)

	// setup (indices, expiration, ...)
	setup()
}

// Creates indices for database and required data.
func setup() {
	log.Print("Setting up database")

	// ...
}

// Disconnect from database. Call on shutdown!
func Disconnect() {
	session.Close()
}

// Returns the database session.
func Get() *mgo.Session {
	return session
}
