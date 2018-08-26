package langmap

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WordService struct {
	Engine *Engine
}

func (s *WordService) Create(c *gin.Context) {
	w := Word{}

	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO(dario) check for errors
	s.Engine.DB.Create(&w)

	c.JSON(http.StatusCreated, nil)
}

func (s *WordService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := Word{}
	w.ID = uint(id)

	// TODO(dario) check for errors
	// this won't delete definitions. should it?
	// it probably should.
	s.Engine.DB.Delete(&w)

	c.JSON(http.StatusNoContent, nil)
}

func (s *WordService) Get(c *gin.Context) {
	w := make([]Word, 0)
	db := s.Engine.DB

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definitions")
	}

	// TODO(dario) check for errors
	db.Find(&w)

	c.JSON(http.StatusOK, w)
	return
}

func (s *WordService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := Word{}
	db := s.Engine.DB

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definitions")
	}

	// TODO(dario) check for errors
	db.Find(&w, id)

	c.JSON(http.StatusOK, w)
	return
}

func (s *WordService) Update(c *gin.Context) {
}

func (s *WordService) Templates() []string {
	return []string{
		"word/new",
	}
}
