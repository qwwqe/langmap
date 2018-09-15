package langmap

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Engine struct {
	Config *Config
	DB     *gorm.DB
	Router *gin.Engine
}

func (e *Engine) AddService(s Service) {
	g := e.Router.Group(s.GetPrefix())

	g.POST("/", s.Create)
	g.DELETE("/:id", s.Delete)
	g.GET("/", s.Get)
	g.GET("/:id", s.GetOne)
	g.PATCH("/:id", s.Update)

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

	if e.Config.Database.Driver == "sqlite3" {
		e.DB.Exec("PRAGMA foreign_keys = ON")
	}

	e.DB.AutoMigrate(&Association{})
	e.DB.AutoMigrate(&Collection{})
	e.DB.AutoMigrate(&Definition{})
	e.DB.AutoMigrate(&Note{})
	e.DB.AutoMigrate(&Tag{})
	e.DB.AutoMigrate(&Usage{})
	e.DB.AutoMigrate(&Word{})

	e.DB.Model(&Word{}).AddUniqueIndex(
		"idx_word_language_user",
		"word",
		"language_code",
		"user_id",
	)

	e.DB.LogMode(e.Config.Database.LogMode)

	return nil
}

func (e *Engine) SetupRouter() {
	e.Router = gin.Default()

	v := VersionMiddleware{Engine: e}
	a := AuthMiddleware{Engine: e}
	e.Router.Use(v.Handler, a.Handler)

	e.Router.HTMLRender = multitemplate.New()

	e.AddService(&DefinitionService{Engine: e, Prefix: "/api/definitions"})
	e.AddService(&NoteService{Engine: e, Prefix: "/api/notes"})
	e.AddService(&UsageService{Engine: e, Prefix: "/api/usages"})
	e.AddService(&WordService{Engine: e, Prefix: "/api/words"})
}

func (e *Engine) Run() error {
	err := e.SetupDB()
	if err != nil {
		return err
	}
	defer e.DB.Close()

	e.SetupRouter()

	e.Router.Run(e.Config.Address)

	return nil
}
