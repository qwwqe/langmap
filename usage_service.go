package langmap

import (
	"database/sql"
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
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := s.Engine.DbMap.Insert(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(u.Id), 10)))

	c.Status(http.StatusCreated)
}

func (s *UsageService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := Usage{}

	if err := s.Engine.DbMap.SelectOne(&u, "select id from usages where id = $1", id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"reason": ErrDatabaseNotFound,
				"errors": NewErrorsJSON([]error{err}),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if _, err := s.Engine.DbMap.Delete(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *UsageService) Get(c *gin.Context) {
	var u []Usage

	if _, err := s.Engine.DbMap.Select(&u, "select * from usages"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}

func (s *UsageService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := Usage{}

	if err := s.Engine.DbMap.SelectOne(&u, "select * from usages where id = $1", id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"reason": ErrDatabaseNotFound,
				"errors": NewErrorsJSON([]error{err}),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}

func (s *UsageService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	u := &Usage{Id: uint(id)}

	if err := s.Engine.DbMap.SelectOne(&u, "select * from usages where id = $1", id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"reason": ErrDatabaseNotFound,
				"errors": NewErrorsJSON([]error{err}),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	u.FromMap(data)

	if _, err := s.Engine.DbMap.Update(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
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
