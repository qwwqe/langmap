package langmap

import (
	"github.com/gin-gonic/gin"
)

type LanguageService struct {
	BaseService
}

func (s *LanguageService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Language{}, c)
}

func (s *LanguageService) Delete(c *gin.Context) {
	ServiceDelete(s, &Language{}, c)
}

func (s *LanguageService) Get(c *gin.Context) {
	ServiceGet(s, Language{}, &[]Language{}, c)
}

func (s *LanguageService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Language{}, c)
}

func (s *LanguageService) Update(c *gin.Context) {
	ServiceUpdate(s, &Language{}, c)
}

func (s *LanguageService) Register() {
	ServiceRegister(s)
}
