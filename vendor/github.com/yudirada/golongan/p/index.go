package p

import "github.com/gin-gonic/gin"

// GenPageIndex contains router for index
func GenPageIndex(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("TEST2"))
	})
}
