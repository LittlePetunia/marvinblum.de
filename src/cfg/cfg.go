package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Host      string `json:"host"`       // e.g. localhost:1234
	DbHost    string `json:"dbhost"`     // MongoDB host with port
	Db        string `json:"db"`         // MongoDB database to use
	DbUser    string `json:"dbuser"`     // MongoDB database user
	DbPwd     string `json:"dbpwd"`      // MongoDB database user password
	LogFile   string `json:"logfile"`    // optional
	Login     string `json:"login"`      // login to page
	PwdSha256 string `json:"pwd_sha256"` // SHA256 password for login
}

// Loads config from json.
func Load(file string) Config {
	config := Config{}

	// read
	log.Print("Loading config file: ", file)
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	// parse
	if err := json.Unmarshal(content, &config); err != nil {
		panic(err)
	}

	return config
}
