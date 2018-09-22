package langmap

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

var (
	Version = "0.0.0"
)

type MapInjectable interface {
	FromMap(map[string]interface{})
}

type IdentifiableResource interface {
	GetId() uint
}

type TableWriter interface {
	Db() *gorp.DbMap
	Table() *gorp.TableMap
	TableName() string
}

type RoutableService interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Get(*gin.Context)
	GetOne(*gin.Context)
	Update(*gin.Context)

	SetEngine(*Engine)
	Router() *gin.RouterGroup
}
