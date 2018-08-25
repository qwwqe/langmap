package langmap

import "github.com/gin-gonic/gin"

type VersionMiddleware struct {
	Engine *Engine
}

func (m *VersionMiddleware) Handler(c *gin.Context) {
	c.Writer.Header().Set("Server", "langmap/"+Version)
	c.Next()
}
