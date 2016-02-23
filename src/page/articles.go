package page

import (
	"html/template"
	"log"
	"net/http"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(page_template_file, bar_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := Page{"", template.HTML("")}
	err = tpl.Execute(w, page)

	if err != nil {
		log.Print(err)
	}
}
