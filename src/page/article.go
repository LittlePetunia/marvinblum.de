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

type removeComment struct {
	Article string    `json:"article"` // article ID
	Created time.Time `json:"created"`
}

type addArticle struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Picture  string `json:"picture"`
	Headline string `json:"headline"`
}

type saveArticle struct {
	addArticle
	Article string `json:"article"`
	Content string `json:"content"`
}

type response struct {
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

	resp := response{true}

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

	resp := response{blog.RemoveCommentByDate(article.Id, remove.Created)}
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	add := addArticle{}

	if err := decoder.Decode(&add); err != nil {
		log.Print(err)
		return
	}

	resp := response{false}

	if !util.IsEmpty(add.Title) && !util.IsEmpty(add.Link) {
		if util.IsEmpty(add.Picture) {
			add.Picture = ""
		}

		resp.Success = blog.AddArticle(add.Title, add.Link, add.Picture, add.Headline)
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func SaveArticleHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	save := saveArticle{}

	if err := decoder.Decode(&save); err != nil {
		log.Print(err)
		return
	}

	resp := response{false}

	if !util.IsEmpty(save.Title) && !util.IsEmpty(save.Link) {
		if util.IsEmpty(save.Picture) {
			save.Picture = ""
		}

		article := blog.FindArticleById(save.Article)

		if article == nil {
			log.Print("Article not found with ID: ", save.Article)
			return
		}

		article.Title = save.Title
		article.Link = save.Link
		article.Picture = save.Picture
		article.Headline = template.HTML(save.Headline)
		article.Content = template.HTML(save.Content)

		resp.Success = blog.SaveArticle(article)
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
