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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Create(&w); db.Error != nil {
		if db.RecordNotFound() {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   ErrDatabaseNotFound,
				"reasons": NewErrorsJSON(db.GetErrors()),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(w.ID), 10)))

	c.Status(http.StatusCreated)
}

func (s *WordService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	w := Word{}
	w.ID = uint(id)

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Delete(&w); db.Error != nil {
		if db.RecordNotFound() {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   ErrDatabaseNotFound,
				"reasons": NewErrorsJSON(db.GetErrors()),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *WordService) Get(c *gin.Context) {
	w := make([]Word, 0)
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definitions.Usages")
	}

	if db := db.Find(&w); db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.JSON(http.StatusOK, w)
	return
}

func (s *WordService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	w := Word{}
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definitions.Usages")
	}

	if db := s.Engine.DB.Find(&w, id); db.Error != nil {
		if db.RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   ErrDatabaseNotFound,
				"reasons": NewErrorsJSON(db.GetErrors()),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.JSON(http.StatusOK, w)
	return
}

func (s *WordService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	w := Word{}
	w.ID = uint(id)

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Model(&w).Updates(data); db.Error != nil {
		if db.RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   ErrDatabaseNotFound,
				"reasons": NewErrorsJSON(db.GetErrors()),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

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
