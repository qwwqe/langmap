package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Note struct {
	gorm.Model
	Word          string `gorm:"UNIQUE;NOT NULL" form:"word"`
	Pronunciation string `form:"pronunciation"`
	Definition    string `form:"definition"`
	Comment       string `form:"comment"`
}

type Association struct {
	gorm.Model
	Note1ID int `gorm:"UNIQUE_INDEX:word_association"`
	Note1   Note
	Note2ID int `gorm:"UNIQUE_INDEX:word_association"`
	Note2   Note
	Type    int `gorm:"UNIQUE_INDEX:word_association"`
	Comment string
}

type Tag struct {
	gorm.Model
	NoteID int
	Note   Note
	Name   string
}

type Collection struct {
	gorm.Model
	NoteID int
	Note   Note
	Name   string
}

func main() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&Note{})
	db.AutoMigrate(&Association{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Collection{})

	r := gin.Default()

	r.HTMLRender = func() multitemplate.Renderer {
		r := multitemplate.NewRenderer()
		r.AddFromFiles("index", "templates/base.tmpl", "templates/index.tmpl")
		r.AddFromFiles("notes_new", "templates/base.tmpl", "templates/notes_new.tmpl")
		r.AddFromFiles("notes_show", "templates/base.tmpl", "templates/notes_show.tmpl")
		return r
	}()

	r.GET("/", func(c *gin.Context) {
		var results []Note
		q := c.Query("query")
		if q != "" {
			db.Where("word LIKE ?", "%" + q + "%").Find(&results)			
		}
		c.HTML(http.StatusOK, "index", gin.H{
			"results": results,
		})
			
	})
	
	{
		n := r.Group("/notes")

		n.GET("/new", func(c *gin.Context) {
			c.HTML(http.StatusOK, "notes_new", nil)
		})

		n.POST("/new", func(c *gin.Context) {
			var note Note

			err := c.ShouldBind(&note)
			if err != nil {
				c.String(http.StatusInternalServerError, "error!", nil)
				return
			}

			note.DeletedAt = nil

			db.Create(&note)

			c.Redirect(http.StatusSeeOther, fmt.Sprintf("/notes/show/%d", note.ID)) // TODO shit
		})

		n.GET("/show/:id", func(c *gin.Context) {
			var note Note

			db.Where("id = ?", c.Param("id")).First(&note)

			c.HTML(http.StatusOK, "notes_show", gin.H{
				"Note": note,
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
