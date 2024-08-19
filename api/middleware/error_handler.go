package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context){
	c.Next()

	for _, err := range c.Errors{
		switch err.Err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Err.Error()})
		default:
			c.JSON(-1, gin.H{"error": err.Err.Error()})
		}
	}
}
