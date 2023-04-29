package v1

import (
	"RedNews_Server/internal/domain/entity"
	mysql2 "RedNews_Server/pkg/client/mysql"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	id := c.Param("id")

	if articles, err := mysql2.GetArticles(id); err != nil {
		c.JSONP(404, gin.H{"Error": err.Error()})
	} else {
		c.JSONP(200, gin.H{"Articles": articles})
	}
}

func NewArticles(c *gin.Context) {
	var articles entity.Articles
	if err := c.ShouldBindJSON(&articles); err != nil {
		c.JSONP(404, gin.H{"Error": "Not should json"})
		return
	}

	if err := mysql2.NewArticles(articles); err != nil {
		c.JSONP(404, gin.H{"Error": err.Error()})
		return
	}

	c.JSONP(200, gin.H{"message": "articles created!"})
}
