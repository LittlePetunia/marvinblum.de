package page

import (
	"blog"
	"html/template"
)

type Page struct {
	Title       string
	Content     template.HTML
	NewArticles []blog.Article
}

// Creates a new page filled with basic content.
func newPage() *Page {
	page := Page{}
	page.NewArticles = *blog.GetNewArticles(bar_new_article_n)

	return &page
}
