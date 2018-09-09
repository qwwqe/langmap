package langmap

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DefinitionService struct {
	Engine *Engine
	Prefix string
}

func (s *DefinitionService) Create(c *gin.Context) {
	d := Definition{}

	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Create(&d); db.Error != nil {
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

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(d.ID), 10)))

	c.Status(http.StatusCreated)
}

func (s *DefinitionService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := Definition{}
	d.ID = uint(id)

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Delete(&d); db.Error != nil {
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

func (s *DefinitionService) Get(c *gin.Context) {
	d := make([]Definition, 0)
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Usages").Preload("Word")
	}

	if db := db.Find(&d); db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.JSON(http.StatusOK, d)
	return
}

func (s *DefinitionService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := Definition{}
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Usages").Preload("Word")
	}

	if db := s.Engine.DB.Find(&d, id); db.Error != nil {
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

	c.JSON(http.StatusOK, d)
	return
}

func (s *DefinitionService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := Definition{}
	d.ID = uint(id)

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Model(&d).Updates(data); db.Error != nil {
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

func (s *DefinitionService) GetPrefix() string {
	return s.Prefix
}

func (s *DefinitionService) Templates() []string {
	return []string{}
}
