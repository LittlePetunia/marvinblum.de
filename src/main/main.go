package main

import (
	"cfg"
	"db"
	"log"
	"net/http"
	"os"
	"page"
	"util"
)

const (
	config_file = "config.json"
	public_dir  = "public"
)

var (
	config cfg.Config
)

// Log to file if logfile name is set.
func logToFile() *os.File {
	if config.LogFile == "" {
		return nil
	}

	handle, err := os.Create(config.LogFile)

	if err != nil {
		panic(err)
	}

	log.SetOutput(handle)

	return handle
}

// Starts the REST server.
func startServer() {
	log.Print("Starting server on ", config.Host)

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.Dir(public_dir)))
	mux.HandleFunc("/", page.HomeHandler)

	mux.HandleFunc("/article/", page.ArticleHandler)
	mux.HandleFunc("/articles", page.ArticlesHandler)
	mux.Handle("/addArticle", http.HandlerFunc(page.AddArticleHandler))
	mux.Handle("/saveArticle", http.HandlerFunc(page.SaveArticleHandler))
	mux.Handle("/removeArticle", http.HandlerFunc(page.RemoveArticleHandler))
	mux.HandleFunc("/search", page.SearchArticleHandler)
	mux.HandleFunc("/addComment", page.AddCommentHandler)
	mux.Handle("/removeComment", http.HandlerFunc(page.RemoveCommentHandler))
	mux.HandleFunc("/login", page.LoginHandler)
	mux.Handle("/logout", http.HandlerFunc(page.LogoutHandler))
	mux.Handle("/upload", http.HandlerFunc(page.UploadHandler))

	if err := http.ListenAndServe(config.Host, mux); err != nil {
		panic(err)
	}
}

func main() {
	config = cfg.Load(config_file)
	log := logToFile()

	if log != nil {
		defer log.Close()
	}

	db.Connect(config.DbHost, config.Db, config.DbUser, config.DbPwd)
	defer db.Disconnect()
	util.StartSessionManager()
	startServer()
}
