package page

import (
	"blog"
	"html/template"
	"net/http"
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

	// check logged in
	cookie, err := r.Cookie(cookie_name)

	if err == nil {
		token := cookie.Value
		page.LoggedIn = userSession.LoggedIn(token)
	}

	return &page
}
