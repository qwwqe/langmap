package langmap

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WordService struct {
	Engine *Engine
	Prefix string
}

func (s *WordService) Create(c *gin.Context) {
	w := Word{}

	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO(dario) check for errors
	// success   = http.StatusCreated
	// not found = http.StatusBadRequest
	// other     = http.StatusInternalServer
	s.Engine.DB.Create(&w)

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(w.ID), 10)))

	c.Status(http.StatusCreated)
}

func (s *WordService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := Word{}
	w.ID = uint(id)

	// TODO(dario) this won't delete definitions. should it? it probably should.
	// TODO(dario) check for errors
	// success   = http.StatusNoContent
	// not found = http.StatusBadRequest
	// other     = http.StatusInternalServer
	s.Engine.DB.Delete(&w)

	c.Status(http.StatusNoContent)
}

func (s *WordService) Get(c *gin.Context) {
	w := make([]Word, 0)
	db := s.Engine.DB

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definitions")
	}

	// TODO(dario) check for errors
	// success   = http.StatusOK
	// not found = http.StatusOK
	// other     = http.StatusInternalServer
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
	// success   = http.StatusOK
	// not found = http.StatusNotFound
	// other     = http.StatusInternalServer
	db.Find(&w, id)

	c.JSON(http.StatusOK, w)
	return
}

func (s *WordService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w := Word{}
	w.ID = uint(id)

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO(dario) check for errors
	// success   = http.StatusNoContent
	// not found = http.StatusBadRequest
	// other     = http.StatusInternalServer
	s.Engine.DB.Model(&w).Updates(data)

	c.Status(http.StatusNoContent)
}

func (s *WordService) GetPrefix() string {
	return s.Prefix
}

func (s *WordService) Templates() []string {
	return []string{
		"word/new",
	}
}
