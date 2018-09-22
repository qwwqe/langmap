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

func (s *BaseService) Db() *gorp.DbMap       { return s.Engine.DbMap }
func (s *BaseService) Table() *gorp.TableMap { return s.TableMap }
func (s *BaseService) TableName() string     { return s.TableMap.TableName }

func (s *BaseService) SetEngine(e *Engine) { s.Engine = e }
func (s *BaseService) Router() *gin.RouterGroup {
	g := s.Engine.Router.Group(s.Prefix)

	g.POST("/", s.Create)
	g.DELETE("/:id", s.Delete)
	g.GET("/", s.Get)
	g.GET("/:id", s.GetOne)
	g.PATCH("/:id", s.Update)

	return g
}
