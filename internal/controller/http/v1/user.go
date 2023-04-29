package v1

import (
	"RedNews_Server/internal/domain/entity"
	mysql2 "RedNews_Server/pkg/client/mysql"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	var users entity.User
	var err error

	id := c.Param("id")
	users, err = mysql2.GetUser(id)

	if err != nil {
		c.JSONP(404, gin.H{"Error": err})
		return
	}

	c.JSONP(202, gin.H{"User": users})
	return
}

func Registration(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8081/reg")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSONP(404, gin.H{"Error": "Not found field"})
		return
	}

	if err := mysql2.Register(&user); err != nil {
		c.JSONP(404, gin.H{"Error": err.Error()})
		return
	}

	c.JSONP(200, nil)
}

func Authorization(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSONP(404, gin.H{"Error": "Not found field"})
		return
	}

	if user, err := mysql2.Authorization(&user); err == nil {
		c.JSONP(200, gin.H{"User": user})
	} else {
		c.JSONP(404, gin.H{"Error": err})
	}
}
