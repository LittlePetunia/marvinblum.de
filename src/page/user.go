package page

import (
	"cfg"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"session"
	"strings"
	"sync"
)

const (
	config_file         = "config.json"
	login_template_file = "public/tpl/login.html"
	login_content_title = "marvinblum.de - Login"

	cookie_name  = "mb_login"
	max_lifetime = 14400 // 4 hours
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

var (
	m           sync.Mutex
	userSession session.Session
)

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
		m.Lock()
		userSession.Login(login.Login, login.Password)
		m.Unlock()

		resp.Success = true

		// create cookie
		cookie, _ := r.Cookie(cookie_name)
		cookie = &http.Cookie{Name: cookie_name,
			Value:    userSession.GetToken(),
			MaxAge:   max_lifetime,
			HttpOnly: true}
		http.SetCookie(w, cookie)
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	userSession.Logout()
	m.Unlock()
	http.Redirect(w, r, "/", 301)
}

// Middleware to check user session.
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check cookie
		cookie, err := r.Cookie(cookie_name)

		if err != nil {
			w.Write([]byte("Not logged in!"))
			return
		}

		token := cookie.Value

		if !userSession.LoggedIn(token) {
			w.Write([]byte("Not logged in!"))
			return
		}

		// go on
		next.ServeHTTP(w, r)
	})
}
