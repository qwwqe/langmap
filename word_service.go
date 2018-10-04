package langmap

import (
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
	ServiceGet(s, Word{}, &[]Word{}, c)
}

func (s *WordService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Word{}, c)
}

func (s *WordService) Update(c *gin.Context) {
	ServiceUpdate(s, &Word{}, c)
}
