package http

import (
	v1 "RedNews_Server/internal/controller/http/v1"
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	user := server.Group("/user")
	{
		user.POST("/reg", v1.Registration)
		user.POST("/auth", v1.Authorization)
		user.POST("/:id", v1.GetUser)
	}

	articles := server.Group("/articles")
	{
		articles.POST("/:id", v1.GetArticles)
		articles.POST("/new", v1.NewArticles)
	}
}
