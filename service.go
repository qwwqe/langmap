package langmap

import "github.com/gin-gonic/gin"

type Service interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Get(*gin.Context)
	GetOne(*gin.Context)
	Update(*gin.Context)
	Templates() []string
}
