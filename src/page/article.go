package page

import (
	"blog"
	"encoding/json"
	"html"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"util"
)

const (
	article_template_file = "public/tpl/article.html"
)

type articlePage struct {
	page
	Article blog.Article
}

type saveComment struct {
	Article string `json:"article"` // article ID
	Name    string `json:"name"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

type saveCommentResponse struct {
	Success bool `json:"success"`
}

type removeComment struct {
	Article string    `json:"article"` // article ID
	Created time.Time `json:"created"`
}

type removeCommentResponse struct {
	Success bool `json:"success"`
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(article_template_file, head_template_file, foot_template_file)

	if err != nil {
		log.Print(err)
		return
	}

	link := strings.Split(r.URL.Path, "/")

	if len(link) < 3 {
		log.Print("Could not parse article path")
		return
	}

	article := blog.FindArticleByLink(link[2])

	if article == nil {
		article = &blog.Article{}
	}

	page := newPage()
	pageWithArticle := articlePage{*page, *article}
	err = tpl.Execute(w, pageWithArticle)

	if err != nil {
		log.Print(err)
	}
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	save := saveComment{}

	if err := decoder.Decode(&save); err != nil {
		log.Print(err)
		return
	}

	resp := saveCommentResponse{true}

	// TODO check email
	if util.IsEmpty(save.Name) || util.IsEmpty(save.Email) || util.IsEmpty(save.Comment) {
		resp.Success = false
	}

	if resp.Success {
		article := blog.FindArticleById(save.Article)

		if article == nil {
			resp.Success = false
		} else {
			save.Name = html.EscapeString(save.Name)
			save.Email = html.EscapeString(save.Email)
			save.Comment = strings.Replace(html.EscapeString(save.Comment), "\n", "<br />", -1)
			resp.Success = blog.AddComment(article.Id, save.Name, save.Email, save.Comment)
		}
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func RemoveCommentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	remove := removeComment{}

	if err := decoder.Decode(&remove); err != nil {
		log.Print(err)
		return
	}

	article := blog.FindArticleById(remove.Article)

	if article == nil {
		log.Print("Article not found with ID: ", remove.Article)
		return
	}

	resp := removeCommentResponse{blog.RemoveCommentByDate(article.Id, remove.Created)}
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
