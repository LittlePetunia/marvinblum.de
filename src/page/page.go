package page

import (
	"blog"
	"html/template"
)

type page struct {
	Title       string
	Content     template.HTML
	NewArticles []blog.Article
}

// Creates a new page filled with basic content.
func newPage() *page {
	page := page{}
	page.NewArticles = *blog.GetArticles(bar_new_article_n)

	return &page
}
