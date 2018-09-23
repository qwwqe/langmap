package langmap

import (
	"github.com/gin-gonic/gin"
)

type InstanceService struct {
	BaseService
}

func (s *InstanceService) Create(c *gin.Context) {
	ServiceCreate(s, s.Prefix, &Instance{}, c)
}

func (s *InstanceService) Delete(c *gin.Context) {
	ServiceDelete(s, &Instance{}, c)
}

func (s *InstanceService) Get(c *gin.Context) {
	ServiceGet(s, Instance{}, &[]Instance{}, c)
}

func (s *InstanceService) GetOne(c *gin.Context) {
	ServiceGetOne(s, &Instance{}, c)
}

func (s *InstanceService) Update(c *gin.Context) {
	ServiceUpdate(s, &Instance{}, c)
}

func (s *InstanceService) Register() {
	ServiceRegister(s)
}
