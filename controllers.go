package langmap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IndexController struct {
	DB *gorm.DB
}

func (i IndexController) Index(c *gin.Context) {
	var results []Note
	q := c.Query("query")
	if q != "" {
		i.DB.Where("word LIKE ?", "%"+q+"%").Find(&results)
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"results": results,
	})

}

type NotesController struct {
	DB *gorm.DB
}

func (i NotesController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "notes_new", nil)
}

func (i NotesController) Create(c *gin.Context) {
	var f NoteForm

	err := c.ShouldBind(&f)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	note := Note{}

	i.DB.Create(&note)

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/notes/show/%d", note.ID)) // TODO shit
}

func (i NotesController) Show(c *gin.Context) {
	var note Note

	i.DB.Where("id = ?", c.Param("id")).First(&note)

	c.HTML(http.StatusOK, "notes_show", gin.H{
		"Note": note,
	})
}

type WordsController struct {
	DB *gorm.DB
}

func (i WordsController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "words_new", nil)
}

func (i WordsController) Create(c *gin.Context) {
	type Form struct {
		Meaning       map[string]string `form:"meaning"`
		Pronunciation map[string]string `form:"pronunciation"`
		Word          string            `form:"word"`
	}

	f := Form{}

	err := c.ShouldBind(&f)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("%+v\n", f)

	c.String(http.StatusOK, "test", nil)
}
