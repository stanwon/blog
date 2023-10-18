package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
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
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articles": articles,
	})
}

func run() {
	router := gin.Default()
	router.Static("/css", "static")
	router.LoadHTMLGlob("static/*")

	router.GET("/", index)
	router.GET("/article", getArticles)
	router.GET("/test", test)

	log.Fatal(router.Run(":8080"))
}
func test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}
