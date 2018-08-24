package langmap

import "github.com/gin-gonic/gin"

type NoteService struct {
	Engine *Engine
}

func (s *NoteService) Create(c *gin.Context) {
}

func (s *NoteService) Delete(c *gin.Context) {
}

func (s *NoteService) Get(c *gin.Context) {
}

func (s *NoteService) GetOne(c *gin.Context) {
}

func (s *NoteService) Update(c *gin.Context) {
}

func (s *NoteService) Templates() []string {
	return []string{
		"note/new",
		"note/show",
	}
}