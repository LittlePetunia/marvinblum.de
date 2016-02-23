package page

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	about_content_title = "marvinblum.de - Über mich"
	about_content_file  = "public/tpl/about.html"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(page_template_file, bar_template_file)

	if err != nil {
		log.Fatal(err)
		return
	}

	content, err := ioutil.ReadFile(about_content_file)

	if err != nil {
		log.Fatal(err)
		return
	}

	page := Page{about_content_title, template.HTML(content)}
	err = tpl.Execute(w, page)

	if err != nil {
		log.Fatal(err)
	}
}
