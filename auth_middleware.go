package langmap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Engine *Engine
}

func (m *AuthMiddleware) Handler(c *gin.Context) {
	_, user := c.Request.Header["User-Id"]
	_, lang := c.Request.Header["Language-Code"]

	if !user || !lang {
		c.Status(http.StatusBadRequest)
		c.Abort()
	}
}
