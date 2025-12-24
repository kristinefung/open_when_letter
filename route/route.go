package route

import (
	"github.com/gin-gonic/gin"

	"open_when_letter/internal/handler"
)

func SetupRouter(boxHandler *handler.LetterBoxHandler) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/letterbox", boxHandler.CreateLetterBox)
		// api.GET("/invite/:code", boxHandler.GetByInviteCode)
	}

	return r
}
