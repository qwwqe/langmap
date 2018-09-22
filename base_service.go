package langmap

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

type BaseService struct {
	Engine   *Engine
	Prefix   string
	TableMap *gorp.TableMap
}

func (s *BaseService) Create(c *gin.Context) {}
func (s *BaseService) Delete(c *gin.Context) {}
func (s *BaseService) Get(c *gin.Context)    {}
func (s *BaseService) GetOne(c *gin.Context) {}
func (s *BaseService) Update(c *gin.Context) {}

func (s *BaseService) Db() *gorp.DbMap                       { return s.Engine.DbMap }
func (s *BaseService) Router(r *gin.Engine) *gin.RouterGroup { return r.Group(s.Prefix) }
func (s *BaseService) Table() *gorp.TableMap                 { return s.TableMap }
func (s *BaseService) TableName() string                     { return s.TableMap.TableName }
