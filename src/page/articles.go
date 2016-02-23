package page

import (
	"html/template"
	"log"
	"net/http"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(page_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := newPage()
	err = tpl.Execute(w, page)

	if err != nil {
		log.Print(err)
	}
}
