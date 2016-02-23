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

// Returns the n newest articles or all if n is <= 0.
func GetArticles(n int) *[]Article {
	if n < 0 {
		n = 0
	}

	articles := make([]Article, n)
	var err error

	if n > 0 {
		err = db.Get().C("article").Find(bson.M{}).Sort("created").Limit(n).All(&articles)
	} else {
		err = db.Get().C("article").Find(bson.M{}).Sort("created").All(&articles)
	}

	if err != nil {
		log.Print(err)
		return nil
	}

	return &articles
}

// Finds an article by its normalized title (link).
func FindArticleByLink(link string) *Article {
	var article Article
	err := db.Get().C("article").Find(bson.M{"link": link}).One(&article)

	if err != nil {
		log.Print(err)
		return nil
	}

	return &article
}
