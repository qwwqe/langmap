package langmap

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServiceCreate(w TableWriter, prefix string, r IdentifiableResource, c *gin.Context) {
	if err := c.ShouldBindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := w.Db().Insert(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Writer.Header().Set(
		"Location",
		filepath.Join(
			prefix,
			strconv.FormatInt(int64(r.GetId()), 10),
		),
	)

	c.Status(http.StatusCreated)
}

func ServiceDelete(w TableWriter, r interface{}, c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := w.Db().SelectOne(r, "select id from "+w.TableName()+" where id = $1", id); err != nil {
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

	if _, err := w.Db().Delete(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func ServiceGet(w TableWriter, r interface{}, c *gin.Context) {
	if _, err := w.Db().Select(r, "select * from "+w.TableName()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": r})
}

func ServiceGetOne(w TableWriter, r interface{}, c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := w.Db().SelectOne(r, "select * from "+w.TableName()+" where id = $1", id); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"data": r})
}

func ServiceUpdate(w TableWriter, r MapInjectable, c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrInvalidResourceId,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := w.Db().SelectOne(r, "select * from "+w.TableName()+" where id = $1", id); err != nil {
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

	var d map[string]interface{}

	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	r.FromMap(d)

	if _, err := w.Db().Update(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
