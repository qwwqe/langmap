package langmap

import (
	"database/sql"
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
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := s.Engine.DbMap.Insert(&d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(d.Id), 10)))

	c.Status(http.StatusCreated)
}

func (s *DefinitionService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := Definition{}

	if err := s.Engine.DbMap.SelectOne(&d, "select id from definitions where id = $1", id); err != nil {
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

	if _, err := s.Engine.DbMap.Delete(&d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *DefinitionService) Get(c *gin.Context) {
	var d []Definition

	if _, err := s.Engine.DbMap.Select(&d, "select * from definitions"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": d})
}

func (s *DefinitionService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := Definition{}

	if err := s.Engine.DbMap.SelectOne(&d, "select * from definitions where id = $1", id); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"data": d})
}

func (s *DefinitionService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	d := &Definition{Id: uint(id)}

	if err := s.Engine.DbMap.SelectOne(d, "select * from definitions where id = $1", id); err != nil {
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

	d.FromMap(data)

	if _, err := s.Engine.DbMap.Update(d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
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
