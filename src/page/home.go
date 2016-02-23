package page

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	page_template_file = "public/tpl/page.html"
	bar_template_file  = "public/tpl/bar.html"
	home_content_file  = "public/tpl/home.html"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(page_template_file, bar_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	content, err := ioutil.ReadFile(home_content_file)

	if err != nil {
		log.Fatal(err)
		return
	}

	page := Page{"", template.HTML(content)}
	err = tpl.Execute(w, page)

	if err != nil {
		log.Print(err)
	}
}
