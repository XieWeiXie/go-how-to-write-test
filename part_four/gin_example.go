package partfour

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/name/:id", GetNameHandler)
}

func GetNameHandler(c *gin.Context) {
	id := c.Param("id")
	var record Item
	if dbError := POSTGRES.Where("id = ?", id).First(&record).Error; dbError != nil {
		return
	}
	c.JSON(http.StatusOK, record)
}
