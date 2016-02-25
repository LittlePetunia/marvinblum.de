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
	search_template_file   = "public/tpl/search.html"
	search_content_title   = "marvinblum.de - Suche"
)

type articlesPage struct {
	page
	Articles []blog.Article
}

type articleSearchPage struct {
	articlesPage
	Search string
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

func SearchArticleHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(search_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	search := r.FormValue("search")
	page := newPage()
	page.Title = search_content_title
	pageWithArticles := articlesPage{*page, *blog.SearchArticles(search)}
	pageWithArticlesAndSearch := articleSearchPage{pageWithArticles, search}
	err = tpl.Execute(w, pageWithArticlesAndSearch)

	if err != nil {
		log.Print(err)
	}
}
