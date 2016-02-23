package page

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	head_template_file = "public/tpl/head.html"
	foot_template_file = "public/tpl/foot.html"
	page_template_file = "public/tpl/page.html"
	home_content_file  = "public/tpl/home.html"
	bar_new_article_n  = 5
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(page_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	content, err := ioutil.ReadFile(home_content_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := newPage()
	page.Content = template.HTML(content)
	err = tpl.Execute(w, page)

	if err != nil {
		log.Print(err)
	}
}
