package langmap

import (
	"github.com/gin-gonic/gin"
)

type WordService struct {
	Engine *Engine
}

func (s *WordService) Create(c *gin.Context) {
}

func (s *WordService) Delete(c *gin.Context) {
}

func (s *WordService) Get(c *gin.Context) {
}

func (s *WordService) GetOne(c *gin.Context) {
}

func (s *WordService) Update(c *gin.Context) {
}

func (s *WordService) Templates() []string {
	return []string{
		"word/new",
	}
}
