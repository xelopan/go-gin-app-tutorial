package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// showIndexPage is handler for /
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
			// c.HTML(
			// 	http.StatusOK,
			// 	"article.html",
			// 	gin.H{
			// 		"title":   article.Title,
			// 		"payload": article,
			// 	},
			// )
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}
