package langmap

import (
	"github.com/gin-gonic/gin"
)

type UserService struct {
	BaseService
}

func (s *UserService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &User{}, c)
}

func (s *UserService) Delete(c *gin.Context) {
	ServiceDelete(s, &User{}, c)
}

func (s *UserService) Get(c *gin.Context) {
	ServiceGet(s, User{}, &[]User{}, c)
}

func (s *UserService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &User{}, c)
}

func (s *UserService) Update(c *gin.Context) {
	ServiceUpdate(s, &User{}, c)
}
