package langmap

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

type LanguageService struct {
	Engine *Engine
	Prefix string
	Table  *gorp.TableMap
}

func (s *LanguageService) Create(c *gin.Context) {
	l := Language{}

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := s.Engine.DbMap.Insert(&l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Writer.Header().Set("Location", filepath.Join(s.Prefix, strconv.FormatInt(int64(l.Id), 10)))

	c.Status(http.StatusCreated)
}

func (s *LanguageService) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	l := Language{}

	if err := s.Engine.DbMap.SelectOne(&l, "select id from "+s.Table.TableName+" where id = $1", id); err != nil {
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

	if _, err := s.Engine.DbMap.Delete(&l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *LanguageService) Get(c *gin.Context) {
	var l []Language

	if _, err := s.Engine.DbMap.Select(&l, "select * from "+s.Table.TableName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": l})
}

func (s *LanguageService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	l := Language{}

	if err := s.Engine.DbMap.SelectOne(&l, "select * from "+s.Table.TableName+" where id = $1", id); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"data": l})
}

func (s *LanguageService) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	l := &Language{Id: uint(id)}

	if err := s.Engine.DbMap.SelectOne(l, "select * from "+s.Table.TableName+" where id = $1", id); err != nil {
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

	l.FromMap(data)

	if _, err := s.Engine.DbMap.Update(l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *LanguageService) GetPrefix() string {
	return s.Prefix
}

func (s *LanguageService) Templates() []string {
	return []string{}
}
