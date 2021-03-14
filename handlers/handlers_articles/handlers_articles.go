package handlers_articles

import (
	"github.com/TheEgid/articles-site-go/models_articles"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitArticleRoutes() {
	router := gin.Default()
	router.GET("/article/view/:id", viewArticle)
}

func viewArticle(c *gin.Context) {
	id := c.Param("id")
	found, retArticle := models_articles.GetArticleById(id)
	statusCode := http.StatusNotFound
	title := "Article not Found"
	if found {
		statusCode = http.StatusOK
		title = retArticle.Title
	}

	log.Println("viewArticle :: retArticle = ", retArticle)

	c.HTML(statusCode, "view_article.html",
		gin.H{
			"title":   title,
			"found":   found,
			"payload": retArticle,
		},
	)
}
