package page

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"util"
)

const (
	upload_template_file = "public/tpl/upload.html"
	upload_content_title = "marvinblum.de - Upload"

	upload_dir  = "public"
	upload_path = "assets/article_pictures"
	form_file   = "file"
	max_mem     = 20971520 // 20mb
)

type uploadFile struct {
	Name string
	URL  string
}

type filePage struct {
	page
	Files []uploadFile
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if !util.IsLoggedIn(r) {
		log.Print("User not logged in")
		return
	}

	if r.Method == "GET" {
		uploadPage(w, r)
	} else {
		performUpload(w, r)
	}
}

func uploadPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(upload_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	files := make([]uploadFile, 0)
	dir, err := ioutil.ReadDir(upload_dir + "/" + upload_path)

	if err != nil {
		log.Print(err)
	} else {
		for _, file := range dir {
			if !file.IsDir() {
				files = append(files, uploadFile{file.Name(), upload_path + "/" + file.Name()})
			}
		}
	}

	page := newPage(r)
	page.Title = upload_content_title
	pageWithFiles := filePage{*page, files}
	err = tpl.Execute(w, pageWithFiles)

	if err != nil {
		log.Print(err)
	}
}

func performUpload(w http.ResponseWriter, r *http.Request) {
	// get picture
	r.ParseMultipartForm(max_mem)
	file, handler, err := r.FormFile(form_file)

	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/upload", 301)
		return
	}

	defer file.Close()

	// mkdir if required
	err = os.MkdirAll(upload_dir+"/"+upload_path, 0774)

	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/upload", 301)
		return
	}

	// move file
	path := "/" + handler.Filename
	f, err := os.OpenFile(upload_dir+"/"+upload_path+path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/upload", 301)
		return
	}

	defer f.Close()
	io.Copy(f, file)
	http.Redirect(w, r, "/upload", 301)
}
