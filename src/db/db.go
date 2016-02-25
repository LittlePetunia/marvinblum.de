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

	// text search
	weights := make(map[string]int, 2)
	weights["title"] = 3
	weights["headline"] = 2
	weights["content"] = 1
	articleTextSearch := mgo.Index{
		Name:    "article_text_search",
		Key:     []string{"$text:title", "$text:headline", "$text:content"},
		Weights: weights,
		Unique:  false,
		Sparse:  true}

	if err := db.C("article").EnsureIndex(articleTextSearch); err != nil {
		panic(err)
	}
}

// Disconnect from database. Call on shutdown!
func Disconnect() {
	session.Close()
}

// Returns the database.
func Get() *mgo.Database {
	return db
}

// Tests if passed string is a valid ObjectId and can be passed to ObjectHexId.
func IsValidId(id string) bool {
	return len(id) == 24
}
