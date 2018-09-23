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

type Identifiable interface {
	GetId() uint
}

type DatabaseWriter interface {
	Db() *gorp.DbMap
}

type IdentifiableTable interface {
	Identifiable
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
