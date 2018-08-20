package main

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qwwqe/langmap"
)

func main() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&langmap.Note{})
	db.AutoMigrate(&langmap.Association{})
	db.AutoMigrate(&langmap.Tag{})
	db.AutoMigrate(&langmap.Collection{})

	r := gin.Default()

	r.HTMLRender = func() multitemplate.Renderer {
		r := multitemplate.NewRenderer()
		r.AddFromFiles("index", "templates/base.tmpl", "templates/index.tmpl")
		r.AddFromFiles("words_new", "templates/base.tmpl", "templates/words_new.tmpl")
		r.AddFromFiles("notes_new", "templates/base.tmpl", "templates/notes_new.tmpl")
		r.AddFromFiles("notes_show", "templates/base.tmpl", "templates/notes_show.tmpl")
		return r
	}()

	{
		index := langmap.IndexController{DB: db}
		r.GET("/", index.Index)
	}

	{
		words := langmap.WordsController{DB: db}
		w := r.Group("/words")
		w.GET("/new", words.New)
		w.POST("/new", words.Create)
	}

	{
		notes := langmap.NotesController{DB: db}
		n := r.Group("/notes")
		n.GET("/new", notes.New)
		n.POST("/new", notes.Create)
		n.GET("/show/:id", notes.Show)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
