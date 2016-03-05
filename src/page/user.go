package page

import (
	"cfg"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
	"util"
)

const (
	config_file         = "config.json"
	login_template_file = "public/tpl/login.html"
	login_content_title = "marvinblum.de - Login"
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginPage(w, r)
	} else {
		performLogin(w, r)
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(login_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	page := newPage(r)
	page.Title = login_content_title
	err = tpl.Execute(w, page)

	if err != nil {
		log.Print(err)
	}
}

func performLogin(w http.ResponseWriter, r *http.Request) {
	login := loginRequest{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&login); err != nil {
		log.Print(err)

		resp := response{false}
		respJson, _ := json.Marshal(resp)
		w.Write(respJson)

		return
	}

	config := cfg.Load(config_file)
	resp := loginResponse{false, ""}

	if strings.ToLower(login.Login) == config.Login && login.Password == config.PwdSha256 {
		sm := util.GetSessionManager()
		_, err := sm.CreateSession(w, r)

		if err == nil {
			log.Print("User logged in")
			resp.Success = true
		}
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sm := util.GetSessionManager()
	session, err := sm.GetCurrentSession(r)

	if err == nil {
		session.Destroy(w, r)
	}

	http.Redirect(w, r, "/", 301)
}
