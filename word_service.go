package langmap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordService struct {
	Engine *Engine
}

func (s *WordService) Create(c *gin.Context) {
	f := Word{}

	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.Engine.DB.Create(&f)

	c.JSON(http.StatusCreated, nil)
	return
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
