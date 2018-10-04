package langmap

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type WebService struct {
	BaseService
}

func (s *WebService) Register() {
	switch t := s.Engine.Router.HTMLRender.(type) {
	case multitemplate.Renderer:
		t.AddFromFiles("collections_list", "templates/base.tmpl", "templates/collections/list.tmpl")
		t.AddFromFiles("collections_new", "templates/base.tmpl", "templates/collections/new.tmpl")

		t.AddFromFiles("definitions_list", "templates/base.tmpl", "templates/definitions/list.tmpl")
		t.AddFromFiles("definitions_new", "templates/base.tmpl", "templates/definitions/new.tmpl")

		t.AddFromFiles("instances_list", "templates/base.tmpl", "templates/instances/list.tmpl")
		t.AddFromFiles("instances_new", "templates/base.tmpl", "templates/instances/new.tmpl")
		t.AddFromFiles("instances_show", "templates/base.tmpl", "templates/instances/show.tmpl")
	}

	r := s.Engine.Router.Group(s.Prefix)

	r.GET("/", s.Index)

	r.GET("/instances/new", s.NewInstance)
	r.POST("/instances", s.CreateInstance)
	r.GET("/instance/:instance_id", s.ShowInstance)

	r.GET("/instance/:instance_id/definitions", s.ListDefinitions)
	r.POST("/instance/:instance_id/definitions", s.SaveDefinition)
	r.GET("/instance/:instance_id/definitions/new", s.NewDefinition)

	r.GET("/instance/:instance_id/collections", s.ListCollections)
	r.GET("/instance/:instance_id/collections/new", s.NewCollection)
}

func (s *WebService) Index(c *gin.Context) {
	i, _ := LoadInstances(s.Engine.DB)
	c.HTML(http.StatusOK, "instances_list", gin.H{"Instances": i})
}

func (s *WebService) NewInstance(c *gin.Context) {
	users, _ := LoadUsers(s.Engine.DB)
	languages, _ := LoadLanguages(s.Engine.DB)

	c.HTML(http.StatusOK, "instances_new", gin.H{
		"Users":     users,
		"Languages": languages,
	})
}

func (s *WebService) CreateInstance(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (s *WebService) ShowInstance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("instance_id"), 10, 0)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	i := Instance{}
	LoadOne(s.Engine.DB, &i, uint(id))

	collections, _ := LoadCollections(s.Engine.DB, uint(id))
	definitions, _ := LoadDefinitions(s.Engine.DB, uint(id))
	notes, _ := LoadNotes(s.Engine.DB)

	c.HTML(http.StatusOK, "instances_show", gin.H{
		"Collections": collections,
		"Definitions": definitions,
		"Instance":    i,
		"Notes":       notes,
	})
}

func (s *WebService) ListDefinitions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("instance_id"), 10, 0)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	definitions, _ := LoadDefinitions(s.Engine.DB, uint(id))

	log.Println(definitions)
	c.HTML(http.StatusOK, "definitions_list", gin.H{
		"Definitions": definitions,
	})
}

func (s *WebService) SaveDefinition(c *gin.Context) {
	f := struct {
		InstanceId     uint     `form:"instance_id"`
		Meanings       []string `form:"meaning[]"`
		Pronunciations []string `form:"pronunciation[]"`
		Word           string   `form:"word"`
	}{}

	if err := c.ShouldBind(&f); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for i, m := range f.Meanings {
		if err := InsertOne(s.Engine.DB, &Definition{
			InstanceId:    f.InstanceId,
			Meaning:       m,
			Pronunciation: f.Pronunciations[i],
			Word: &Word{
				Word: f.Word,
			},
		}); err != nil {
			log.Println(err)
		}
	}

	c.Status(http.StatusNoContent)
}
func (s *WebService) NewDefinition(c *gin.Context) {
	c.HTML(http.StatusOK, "definitions_new", gin.H{
		"InstanceId": c.Param("instance_id"),
	})
}

func (s *WebService) ListCollections(c *gin.Context) {
	c.HTML(http.StatusOK, "collections_list", gin.H{
		"InstanceId": c.Param("instance_id"),
	})
}

func (s *WebService) NewCollection(c *gin.Context) {
	c.HTML(http.StatusOK, "collections_new", gin.H{
		"InstanceId": c.Param("instance_id"),
	})
}
