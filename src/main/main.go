package main

import (
	"db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"page"
)

const (
	config_file = "config.json"
	public_dir  = "public"
)

type config struct {
	Host      string `json:"host"`       // e.g. localhost:1234
	DbHost    string `json:"dbhost"`     // MongoDB host with port
	Db        string `json:"db"`         // MongoDB database to use
	LogFile   string `json:"logfile"`    // optional
	Login     string `json:"login"`      // login to page
	PwdSha256 string `json:"pwd_sha256"` // SHA256 password for login
}

var (
	cfg config
)

// Loads config from json.
func loadConfig(file string) {
	cfg = config{}

	// read
	log.Print("Loading config file: ", file)
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	// parse
	if err := json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}
}

// Log to file if logfile name is set.
func logToFile() *os.File {
	if cfg.LogFile == "" {
		return nil
	}

	handle, err := os.Create(cfg.LogFile)

	if err != nil {
		panic(err)
	}

	log.SetOutput(handle)

	return handle
}

// Starts the REST server.
func startServer() {
	log.Print("Starting server on ", cfg.Host)

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.Dir(public_dir)))
	mux.HandleFunc("/", page.HomeHandler)
	mux.HandleFunc("/articles", page.ArticlesHandler)
	mux.HandleFunc("/about", page.AboutHandler)

	if err := http.ListenAndServe(cfg.Host, mux); err != nil {
		panic(err)
	}
}

func main() {
	loadConfig(config_file)
	log := logToFile()

	if log != nil {
		defer log.Close()
	}

	db.Connect(cfg.DbHost, cfg.Db)
	defer db.Disconnect()
	startServer()
}
