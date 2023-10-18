package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id       int
	Title    string
	Desc     string
	Content  string
	CreateBy string
	CreateAt int64
	UpdateAt int64
}

func getArticles(c *gin.Context) {
	articles := make([]*Article, 0, 8)
	articles = append(articles, &Article{
		Id:       0,
		Title:    "test",
		Desc:     "desc",
		Content:  "content",
		CreateBy: "stan",
	})
	articles = append(articles, &Article{
		Id:       1,
		Title:    "test1",
		Desc:     "desc2",
		Content:  "content3",
		CreateBy: "stan4",
	})
	articles = append(articles, &Article{
		Id:       1,
		Title:    "test1",
		Desc:     "desc2",
		Content:  "content3",
		CreateBy: "stan4",
	})
	articles = append(articles, &Article{
		Id:       1,
		Title:    "test1",
		Desc:     "desc2",
		Content:  "content3",
		CreateBy: "stan4",
	})
	c.HTML(http.StatusOK, "articles.html", gin.H{
		"articles": articles,
	})
}
