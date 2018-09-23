package langmap

import (
	"github.com/gin-gonic/gin"
)

type CorpusService struct {
	BaseService
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

func (s *CorpusService) Register() {
	ServiceRegister(s)
}
