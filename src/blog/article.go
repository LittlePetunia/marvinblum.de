package blog

import (
	"db"
	"gopkg.in/mgo.v2"
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
	Headline template.HTML `bson:"headline"`
	Content  template.HTML `bson:"content"`
	Comments []Comment     `bson:"comments"`
}

// Returns the n newest articles or all if n is <= 0.
// If full is not set to true, only the id, link, title, headline, picture and created will be returned.
func GetArticles(n int, full bool) *[]Article {
	if n < 0 {
		n = 0
	}

	articles := make([]Article, n)
	var query *mgo.Query

	if full {
		query = db.Get().C("article").Find(bson.M{}).Sort("-created")
	} else {
		query = db.Get().C("article").Find(bson.M{}).Select(bson.M{"_id": 1, "link": 1, "created": 1, "title": 1, "headline": 1, "picture": 1}).Sort("-created")
	}

	if n > 0 {
		query.Limit(n)
	}

	err := query.All(&articles)

	if err != nil {
		log.Print(err)
		return nil
	}

	return &articles
}

// Finds an article by ID.
// The ID will be parsed to ObjectId if valid. nil will be returned on error.
func FindArticleById(id string) *Article {
	if !db.IsValidId(id) {
		log.Print("ID ", id, " is not valid")
		return nil
	}

	var article Article
	err := db.Get().C("article").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&article)

	if err != nil {
		log.Print(err)
		return nil
	}

	return &article
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

// Creates a new article with given title, link and picture.
func AddArticle(title, link, picture, headline string) bool {
	article := Article{Created: time.Now(),
		Updated:  time.Now(),
		Link:     link,
		Title:    title,
		Picture:  picture,
		Headline: template.HTML(headline)}

	err := db.Get().C("article").Insert(article)

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

// Saves and existing article.
func SaveArticle(article *Article) bool {
	article.Updated = time.Now()

	err := db.Get().C("article").Update(bson.M{"_id": article.Id}, article)

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

// Searchs for an article by keywords in its content and title.
func SearchArticles(search string) *[]Article {
	articles := make([]Article, 0)
	err := db.Get().C("article").Pipe([]bson.M{
		bson.M{"$match": bson.M{"$text": bson.M{"$search": search}}},
		bson.M{"$sort": bson.M{"title": 1}}}).All(&articles)

	if err != nil {
		log.Print(err)
		return nil
	}

	return &articles
}

// Removes an article by ID.
func RemoveArticleById(id string) bool {
	if !db.IsValidId(id) {
		log.Print("ID ", id, " is not valid")
		return false
	}

	err := db.Get().C("article").Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}
