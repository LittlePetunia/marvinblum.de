package page

import (
	"blog"
	"html/template"
	"net/http"
	"util"
)

type page struct {
	Title       string
	Content     template.HTML
	NewArticles []blog.Article
	LoggedIn    bool
}

// Creates a new page filled with basic content.
func newPage(r *http.Request) *page {
	page := page{}
	page.NewArticles = *blog.GetArticles(bar_new_article_n, false)
	page.LoggedIn = false

	if util.IsLoggedIn(r) {
		page.LoggedIn = true
	}

	return &page
}
