package blog

import (
	"db"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"time"
)

type Article struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Link     string        `bson:"link"` // build by title, unique
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
	Picture  string        `bson:"picture"`
	Title    string        `bson:"title"`
	Content  template.HTML `bson:"content"`
	Comments []Comment     `bson:"comments"`
}

type Comment struct {
	Created time.Time `bson:"created"`
	Author  string    `bson:"author"`
	Email   string    `bson:"email"`
	Content string    `bson:"content"`
}

// Returns the n newest articles.
func GetNewArticles(n int) *[]Article {
	articles := make([]Article, n)
	err := db.Get().C("article").Find(bson.M{}).Sort("created").Limit(n).All(&articles)

	if err != nil {
		log.Print(err)
		return nil
	}

	return &articles
}
