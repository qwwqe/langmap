package langmap

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

type BaseService struct {
	Engine *Engine
	Prefix string
}

func (s BaseService) Db() *gorp.DbMap { return s.Engine.DbMap }

func (s *BaseService) Create(c *gin.Context) { c.Status(http.StatusNotImplemented) }
func (s *BaseService) Delete(c *gin.Context) { c.Status(http.StatusNotImplemented) }
func (s *BaseService) Get(c *gin.Context)    { c.Status(http.StatusNotImplemented) }
func (s *BaseService) GetOne(c *gin.Context) { c.Status(http.StatusNotImplemented) }
func (s *BaseService) Update(c *gin.Context) { c.Status(http.StatusNotImplemented) }

func (s *BaseService) SetEngine(e *Engine)      { s.Engine = e }
func (s *BaseService) Router() *gin.RouterGroup { return s.Engine.Router.Group(s.Prefix) }

func (s *BaseService) Register() {
	ServiceRegister(s)
}
