package page

import (
	"blog"
	"html/template"
	"log"
	"net/http"
)

const (
	articles_template_file = "public/tpl/articles.html"
	articles_content_title = "marvinblum.de - Artikel"
)

type articlesPage struct {
	page
	Articles []blog.Article
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(articles_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := newPage()
	page.Title = articles_content_title
	pageWithArticles := articlesPage{*page, *blog.GetArticles(0, false)}
	err = tpl.Execute(w, pageWithArticles)

	if err != nil {
		log.Print(err)
	}
}
