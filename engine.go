package langmap

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Engine struct {
	Config *Config
	DB     *gorm.DB
	Router *gin.Engine
}

func (e *Engine) AddService(g *gin.RouterGroup, s Service) {
	g.POST("/", s.Create)
	g.DELETE("/:id", s.Delete)
	g.GET("/", s.Get)
	g.GET("/:id", s.GetOne)
	g.PUT("/:id", s.Update)

	for _, t := range s.Templates() {
		e.Router.HTMLRender.(multitemplate.Render).AddFromFiles(t, "templates/base.tmpl", "templates/"+t+".tmpl")
	}
}

func (e *Engine) SetupDB() error {
	var err error
	e.DB, err = gorm.Open(e.Config.Database.Driver, e.Config.Database.Source)
	if err != nil {
		return err
	}
	defer e.DB.Close()

	e.DB.AutoMigrate(&Association{})
	e.DB.AutoMigrate(&Collection{})
	e.DB.AutoMigrate(&Note{})
	e.DB.AutoMigrate(&Tag{})

	e.DB.LogMode(true)

	return nil
}

func (e *Engine) SetupRouter() {
	e.Router = gin.Default()

	e.Router.HTMLRender = multitemplate.New()

	e.AddService(
		e.Router.Group("/api/words"),
		&WordService{Engine: e},
	)

	e.AddService(
		e.Router.Group("/api/notes"),
		&NoteService{Engine: e},
	)
}

func (e *Engine) Run() error {
	err := e.SetupDB()
	if err != nil {
		return err
	}

	e.SetupRouter()

	e.Router.Run(e.Config.Address)

	return nil
}
