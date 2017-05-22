package p

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderIndex contains router for index
func RenderIndex(r *gin.Engine) {
	r.GET("/", pageIndex)
}

// PageIndex render Index page
func pageIndex(c *gin.Context) {
	//c.Writer.Write([]byte("TEST3"))
	c.HTML(http.StatusOK, "d_login.html", nil)

}
