package langmap

import (
	"github.com/gin-gonic/gin"
)

type DefinitionService struct {
	BaseService
}

func (s *DefinitionService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Definition{}, c)
}

func (s *DefinitionService) Delete(c *gin.Context) {
	ServiceDelete(s, &Definition{}, c)
}

func (s *DefinitionService) Get(c *gin.Context) {
	ServiceGet(s, &[]Definition{}, c)
}

func (s *DefinitionService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Definition{}, c)
}

func (s *DefinitionService) Update(c *gin.Context) {
	ServiceUpdate(s, &Definition{}, c)
}
