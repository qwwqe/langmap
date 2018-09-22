package langmap

import (
	"github.com/gin-gonic/gin"
)

type NoteService struct {
	BaseService
}

func (s *NoteService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Note{}, c)
}

func (s *NoteService) Delete(c *gin.Context) {
	ServiceDelete(s, &Note{}, c)
}

func (s *NoteService) Get(c *gin.Context) {
	ServiceGet(s, &[]Note{}, c)
}

func (s *NoteService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Note{}, c)
}

func (s *NoteService) Update(c *gin.Context) {
	ServiceUpdate(s, &Note{}, c)
}
