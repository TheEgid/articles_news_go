package routes

import (
	"github.com/TheEgid/articles-site-go/handlers/handlers_articles"
	"github.com/TheEgid/articles-site-go/handlers/handlers_users"
	"github.com/TheEgid/articles-site-go/models_articles"
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		gin.H{
			"title":   "Articles",
			"payload": models_articles.GetAllArticles(),
		},
	)
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html",
		gin.H{
			"title": "About Us",
		},
	)
}

func noRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "page_not_found.html",
		gin.H{
			"title": "Page not found",
		},
	)
}

func InitializeRoutes(router *gin.Engine) {
	router.NoRoute(noRouteHandler)
	router.GET("/", homePage)
	router.GET("/about", aboutPage)
	handlers_articles.InitArticleRoutes()
	handlers_users.InitUserRoutes()
}
