package page

import (
	"blog"
	"html/template"
	"log"
	"net/http"
)

const (
	article_template_file = "public/tpl/article.html"
)

type articlePage struct {
	page
	Article blog.Article
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(article_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	article := blog.FindArticleByLink("Mein_erster_Artikel") // TODO

	if article == nil {
		article = &blog.Article{}
	}

	page := newPage()
	pageWithArticle := articlePage{*page, *article}
	err = tpl.Execute(w, pageWithArticle)

	if err != nil {
		log.Print(err)
	}
}
