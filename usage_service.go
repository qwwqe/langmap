package langmap

import (
	"github.com/gin-gonic/gin"
)

type UsageService struct {
	BaseService
}

func (s *UsageService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Usage{}, c)
}

func (s *UsageService) Delete(c *gin.Context) {
	ServiceDelete(s, &Usage{}, c)
}

func (s *UsageService) Get(c *gin.Context) {
	ServiceGet(s, &[]Usage{}, c)
}

func (s *UsageService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Usage{}, c)
}

func (s *UsageService) Update(c *gin.Context) {
	ServiceUpdate(s, &Usage{}, c)
}
