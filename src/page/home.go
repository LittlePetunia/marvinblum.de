package page

import (
	"blog"
	"html/template"
	"log"
	"net/http"
)

const (
	head_template_file = "public/tpl/head.html"
	foot_template_file = "public/tpl/foot.html"
	page_template_file = "public/tpl/page.html"
	home_template_file = "public/tpl/home.html"
	home_new_article_n = 10
	bar_new_article_n  = 5
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(home_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := newPage(r)
	page.Title = articles_content_title
	pageWithArticles := articlesPage{*page, *blog.GetArticles(home_new_article_n, false)}
	err = tpl.Execute(w, pageWithArticles)

	if err != nil {
		log.Print(err)
	}
}
