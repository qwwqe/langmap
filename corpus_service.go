package langmap

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CorpusService struct {
	BaseService
}

func (s *CorpusService) Register() {
	r := s.Engine.Router.Group(s.Prefix)

	RegisterResource(r, s)

	r.POST("/upload", s.Tokenize)
}

func (s *CorpusService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Corpus{}, c)
}

func (s *CorpusService) Delete(c *gin.Context) {
	ServiceDelete(s, &Corpus{}, c)
}

func (s *CorpusService) Get(c *gin.Context) {
	ServiceGet(s, Corpus{}, &[]Corpus{}, c)
}

func (s *CorpusService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Corpus{}, c)
}

func (s *CorpusService) Update(c *gin.Context) {
	ServiceUpdate(s, &Corpus{}, c)
}

func (s *CorpusService) Tokenize(c *gin.Context) {
	r := &struct {
		InstanceId uint `form:"instance_id"`
	}{}

	if err := c.ShouldBind(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	file, err := c.FormFile("text")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "Failed to receive file",
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	i := &Instance{}

	if err := s.Db().SelectOne(i, "select * from "+i.TableName()+" where id = $1", r.InstanceId); err != nil {
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

	if err := i.Preload(s.Db()); err != nil {
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

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Failed to open received file",
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	var t TokenizerAdapter
	switch i.Language.Tag {
	case "ja":
		t = JapaneseTokenizerAdapter{}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "your instance's language does not have a supported tokenizer (available: ja)",
		})
		return

	}

	corpus, err := t.Tokenize(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Failed to open received file",
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	corpus.InstanceId = r.InstanceId

	txn, err := s.Db().Begin()

	if err := txn.Insert(corpus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	if err := txn.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": ErrDatabaseFailure,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	c.Writer.Header().Set(
		"Location",
		filepath.Join(
			"/api",
			corpus.TableName(),
			strconv.FormatUint(uint64(corpus.GetId()), 10),
		),
	)

	c.Status(http.StatusNoContent)
}
