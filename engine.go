package langmap

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

type Engine struct {
	Config *Config
	DbMap  *gorp.DbMap
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
	e.DbMap = &gorp.DbMap{
		Dialect: gorp.SqliteDialect{},
	}

	switch e.Config.Database.Driver {
	case "sqlite3":
		e.DbMap.Dialect = gorp.SqliteDialect{}

	case "postgres":
		e.DbMap.Dialect = gorp.PostgresDialect{}

	}

	var err error
	e.DbMap.Db, err = sql.Open(e.Config.Database.Driver, e.Config.Database.Source)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to open %s database at %s: %s", e.Config.Database.Driver, e.Config.Database.Source, err.Error()))
	}

	var (
		collection_tags       = e.DbMap.AddTableWithName(CollectionTag{}, "collection_tags").SetKeys(true, "Id")
		collections           = e.DbMap.AddTableWithName(Collection{}, "collections").SetKeys(true, "Id")
		corpus_tags           = e.DbMap.AddTableWithName(CorpusTag{}, "corpus_tags").SetKeys(true, "Id")
		corpus_words          = e.DbMap.AddTableWithName(CorpusWord{}, "corpus_words").SetKeys(true, "Id")
		corpora               = e.DbMap.AddTableWithName(Corpus{}, "corpora").SetKeys(true, "Id")
		definition_link_types = e.DbMap.AddTableWithName(DefinitionLinkType{}, "definition_link_types").SetKeys(true, "Id")
		definition_links      = e.DbMap.AddTableWithName(DefinitionLink{}, "definition_links").SetKeys(true, "Id")
		definitions           = e.DbMap.AddTableWithName(Definition{}, "definitions").SetKeys(true, "Id")
		highlights            = e.DbMap.AddTableWithName(Highlight{}, "highlights").SetKeys(true, "Id")
		instances             = e.DbMap.AddTableWithName(Instance{}, "instances").SetKeys(true, "Id")
		languages             = e.DbMap.AddTableWithName(Language{}, "languages").SetKeys(true, "Id")
		lexica                = e.DbMap.AddTableWithName(Lexica{}, "lexica").SetKeys(true, "Id")
		note_collections      = e.DbMap.AddTableWithName(NoteCollection{}, "note_collections").SetKeys(true, "Id")
		note_definitions      = e.DbMap.AddTableWithName(NoteDefinition{}, "note_definitions").SetKeys(true, "Id")
		note_tags             = e.DbMap.AddTableWithName(NoteTag{}, "note_tags").SetKeys(true, "Id")
		notes                 = e.DbMap.AddTableWithName(Note{}, "notes").SetKeys(true, "Id")
		tags                  = e.DbMap.AddTableWithName(Tag{}, "tags").SetKeys(true, "Id")
		usages                = e.DbMap.AddTableWithName(Usage{}, "usages").SetKeys(true, "Id")
		users                 = e.DbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
		wordlist_items        = e.DbMap.AddTableWithName(WordlistItem{}, "wordlist_items").SetKeys(true, "Id")
		wordlists             = e.DbMap.AddTableWithName(Wordlist{}, "wordlist").SetKeys(true, "Id")
		words                 = e.DbMap.AddTableWithName(Word{}, "words").SetKeys(true, "Id")
	)

	if err := e.DbMap.CreateTablesIfNotExists(); err != nil {
		return errors.New("failed to create tables: " + err.Error())
	}

	e.AddForeignKey(collection_tags, "collection_id", collections, "id", 1)
	e.AddForeignKey(collection_tags, "instance_id", instances, "id", 1)
	e.AddForeignKey(collection_tags, "tag_id", tags, "id", 1)

	e.AddForeignKey(collections, "instance_id", instances, "id", 1)

	e.AddForeignKey(corpora, "instance_id", instances, "id", 1)

	e.AddForeignKey(corpus_tags, "corpus_id", corpora, "id", 1)
	e.AddForeignKey(corpus_tags, "instance_id", instances, "id", 1)
	e.AddForeignKey(corpus_tags, "tag_id", tags, "id", 1)

	e.AddForeignKey(corpus_words, "corpus_id", corpora, "id", 1)
	e.AddForeignKey(corpus_words, "word_id", words, "id", 1)

	e.AddForeignKey(definition_links, "definition1_id", definitions, "id", 1)
	e.AddForeignKey(definition_links, "definition2_id", definitions, "id", 2)
	e.AddForeignKey(definition_links, "instance_id", instances, "id", 1)
	e.AddForeignKey(definition_links, "type_id", definition_link_types, "id", 1)

	e.AddForeignKey(definitions, "instance_id", instances, "id", 1)
	e.AddForeignKey(definitions, "word_id", words, "id", 1)

	e.AddForeignKey(highlights, "corpus_id", corpora, "id", 1)
	e.AddForeignKey(highlights, "corpus_word_id", corpus_words, "id", 1)
	e.AddForeignKey(highlights, "instance_id", instances, "id", 1)

	e.AddForeignKey(instances, "language_id", languages, "id", 1)
	e.AddForeignKey(instances, "user_id", users, "id", 1)

	e.AddForeignKey(lexica, "language_id", languages, "id", 1)

	e.AddForeignKey(note_collections, "collection_id", collections, "id", 1)
	e.AddForeignKey(note_collections, "instance_id", instances, "id", 1)
	e.AddForeignKey(note_collections, "note_id", notes, "id", 1)

	e.AddForeignKey(note_definitions, "definition_id", definitions, "id", 1)
	e.AddForeignKey(note_definitions, "instance_id", instances, "id", 1)
	e.AddForeignKey(note_definitions, "note_id", notes, "id", 1)

	e.AddForeignKey(note_tags, "instance_id", instances, "id", 1)
	e.AddForeignKey(note_tags, "note_id", notes, "id", 1)
	e.AddForeignKey(note_tags, "tag_id", tags, "id", 1)

	e.AddForeignKey(notes, "instance_id", instances, "id", 1)

	e.AddForeignKey(tags, "instance_id", instances, "id", 1)

	e.AddForeignKey(usages, "corpus_id", corpora, "id", 1)
	e.AddForeignKey(usages, "definition_id", definitions, "id", 1)
	e.AddForeignKey(usages, "instance_id", instances, "id", 1)

	e.AddForeignKey(wordlist_items, "word_id", words, "id", 1)
	e.AddForeignKey(wordlist_items, "wordlist_id", wordlists, "id", 1)

	e.AddForeignKey(wordlists, "instance_id", instances, "id", 1)

	switch e.Config.Database.Driver {
	case "sqlite3":
		e.DbMap.Exec("PRAGMA foreign_keys = ON")

	}

	if e.Config.Database.LogMode {
		e.DbMap.TraceOn("[gorp]", log.New(os.Stdout, "langsuite:", log.Lmicroseconds))
	}

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
	defer e.DbMap.Db.Close()

	e.SetupRouter()

	e.Router.Run(e.Config.Address)

	return nil
}

// don't like this
// this will produce an error when it already exists but that's ok for now
func (e *Engine) AddForeignKey(table *gorp.TableMap, key string, reference *gorp.TableMap, column string, ordinal uint) error {
	if _, err := e.DbMap.Exec(fmt.Sprintf(
		"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s);",
		e.DbMap.Dialect.QuoteField(table.TableName),
		e.DbMap.Dialect.QuoteField(fmt.Sprintf("fk_%s_%s_%d", table.TableName, reference.TableName, ordinal)),
		e.DbMap.Dialect.QuoteField(key),
		e.DbMap.Dialect.QuoteField(reference.TableName),
		e.DbMap.Dialect.QuoteField(column),
	)); err != nil {
		return err
	}

	return nil
}
