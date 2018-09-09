package langmap

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsageService struct {
	Engine *Engine
	Prefix string
}

func (s *UsageService) Create(c *gin.Context) {
	u := Usage{}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Create(&u); db.Error != nil {
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

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(u.ID), 10)))

	c.Status(http.StatusCreated)
}

func (s *UsageService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := Usage{}
	u.ID = uint(id)

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Delete(&u); db.Error != nil {
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

func (s *UsageService) Get(c *gin.Context) {
	u := make([]Usage, 0)
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definition.Word")
	}

	if db := db.Find(&u); db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   ErrDatabaseFailure,
			"reasons": NewErrorsJSON(db.GetErrors()),
		})
		return
	}

	c.JSON(http.StatusOK, u)
	return
}

func (s *UsageService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := Usage{}
	db := s.Engine.DB.Scopes(UserLanguage(c))

	if _, ok := c.GetQuery("preload"); ok {
		db = db.Preload("Definition.Word")
	}

	if db := s.Engine.DB.Find(&u, id); db.Error != nil {
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

	c.JSON(http.StatusOK, u)
	return
}

func (s *UsageService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrInvalidResourceId,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := Usage{}
	u.ID = uint(id)

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ErrJsonFailed,
			"reasons": NewErrorsJSON([]error{err}),
		})
		return
	}

	db := s.Engine.DB.Scopes(UserLanguage(c))

	if db := db.Model(&u).Updates(data); db.Error != nil {
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

func (s *UsageService) GetPrefix() string {
	return s.Prefix
}

func (s *UsageService) Templates() []string {
	return []string{}
}
