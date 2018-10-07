package langmap

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type WordService struct {
	BaseService
}

func (s *WordService) Register() {
	RegisterResource(s.Engine.Router.Group(s.Prefix), s)
}

func (s *WordService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Word{}, c)
}

func (s *WordService) Delete(c *gin.Context) {
	ServiceDelete(s, &Word{}, c)
}

func (s *WordService) Get(c *gin.Context) {
	r, err := LoadWords(s.Engine.DB, Filter{})
	if err != nil {
		c.JSON(ApiResponseJSON(r, ErrDatabaseFailure, err))
		return
	}
	c.JSON(ApiResponseJSON(r, "", nil))
}

func (s *WordService) GetOne(c *gin.Context) {
	r := Word{}

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(ApiResponseJSON(r, ErrInvalidResourceId, err))
		return
	}

	if err := LoadOne(s.Engine.DB, &r, uint(id)); err != nil {
		c.JSON(ApiResponseJSON(r, ErrDatabaseFailure, err))
		return
	}

	c.JSON(ApiResponseJSON(r, "", nil))
}

func (s *WordService) Update(c *gin.Context) {
	ServiceUpdate(s, &Word{}, c)
}
