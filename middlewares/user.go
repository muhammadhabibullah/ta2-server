package middlewares

import (
	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
)

type User = models.User //

// AuthorizedUser blocks unauthorized requestrs
func AuthorizedUser(c *gin.Context) models.User {
	userRaw, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
	}
	return userRaw.(User)
}
