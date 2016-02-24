package blog

import (
	"db"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"time"
)

type Comment struct {
	Created time.Time     `bson:"created"`
	Author  string        `bson:"author"`
	Email   string        `bson:"email"`
	Content template.HTML `bson:"content"`
}

// Adds a comment to an existing article.
func AddComment(articleId bson.ObjectId, name, email, content string) bool {
	comment := Comment{time.Now(), name, email, template.HTML(content)}

	err := db.Get().C("article").Update(bson.M{"_id": articleId},
		bson.M{"$push": bson.M{"comments": bson.M{"$each": []Comment{comment}, "$position": 0}}})

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

// Removes a comment by its creation date.
func RemoveCommentByDate(articleId bson.ObjectId, date time.Time) bool {
	err := db.Get().C("article").Update(bson.M{"_id": articleId},
		bson.M{"$pull": bson.M{"comments": bson.M{"created": date}}})

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}
