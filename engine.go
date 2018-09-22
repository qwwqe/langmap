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

func (e *Engine) AddService(s RoutableService) {
	g := s.Router(e.Router)

	g.POST("/", s.Create)
	g.DELETE("/:id", s.Delete)
	g.GET("/", s.Get)
	g.GET("/:id", s.GetOne)
	g.PATCH("/:id", s.Update)
}

func (e *Engine) Run() error {

	// Setup DB

	e.DbMap = &gorp.DbMap{
		Dialect: gorp.SqliteDialect{},
	}

	if e.Config.Database.LogMode {
		e.DbMap.TraceOn("[gorp]", log.New(os.Stdout, "langsuite:", log.Lmicroseconds))
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
	defer e.DbMap.Db.Close()

	var (
		collection_tags       = e.DbMap.AddTableWithName(CollectionTag{}, "collection_tags").SetKeys(true, "id")
		collections           = e.DbMap.AddTableWithName(Collection{}, "collections").SetKeys(true, "id")
		corpus_tags           = e.DbMap.AddTableWithName(CorpusTag{}, "corpus_tags").SetKeys(true, "id")
		corpus_words          = e.DbMap.AddTableWithName(CorpusWord{}, "corpus_words").SetKeys(true, "id")
		corpora               = e.DbMap.AddTableWithName(Corpus{}, "corpora").SetKeys(true, "id")
		definition_link_types = e.DbMap.AddTableWithName(DefinitionLinkType{}, "definition_link_types").SetKeys(true, "id")
		definition_links      = e.DbMap.AddTableWithName(DefinitionLink{}, "definition_links").SetKeys(true, "id")
		definitions           = e.DbMap.AddTableWithName(Definition{}, "definitions").SetKeys(true, "id")
		highlights            = e.DbMap.AddTableWithName(Highlight{}, "highlights").SetKeys(true, "id")
		instances             = e.DbMap.AddTableWithName(Instance{}, "instances").SetKeys(true, "id")
		languages             = e.DbMap.AddTableWithName(Language{}, "languages").SetKeys(true, "id")
		lexica                = e.DbMap.AddTableWithName(Lexica{}, "lexica").SetKeys(true, "id")
		note_collections      = e.DbMap.AddTableWithName(NoteCollection{}, "note_collections").SetKeys(true, "id")
		note_definitions      = e.DbMap.AddTableWithName(NoteDefinition{}, "note_definitions").SetKeys(true, "id")
		note_tags             = e.DbMap.AddTableWithName(NoteTag{}, "note_tags").SetKeys(true, "id")
		notes                 = e.DbMap.AddTableWithName(Note{}, "notes").SetKeys(true, "id")
		tags                  = e.DbMap.AddTableWithName(Tag{}, "tags").SetKeys(true, "id")
		usages                = e.DbMap.AddTableWithName(Usage{}, "usages").SetKeys(true, "id")
		users                 = e.DbMap.AddTableWithName(User{}, "users").SetKeys(true, "id")
		wordlist_items        = e.DbMap.AddTableWithName(WordlistItem{}, "wordlist_items").SetKeys(true, "id")
		wordlists             = e.DbMap.AddTableWithName(Wordlist{}, "wordlist").SetKeys(true, "id")
		words                 = e.DbMap.AddTableWithName(Word{}, "words").SetKeys(true, "id")
	)

	words.AddIndex("words_unique_idx", "Btree", []string{"word"}).SetUnique(true)

	if err := e.DbMap.CreateTablesIfNotExists(); err != nil {
		return errors.New("failed to create tables: " + err.Error())
	}

	e.DbMap.CreateIndex()
	// if err := e.DbMap.CreateIndex(); err != nil {
	// 	return errors.New("failed to create indexes: " + err.Error())
	// }

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

	// Setup Router

	e.Router = gin.Default()

	v := VersionMiddleware{Engine: e}

	e.Router.Use(
		v.Handler,
	)

	e.Router.HTMLRender = multitemplate.New()

	e.AddService(&DefinitionService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/definitions",
			TableMap: definitions,
		},
	})

	e.AddService(&NoteService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/notes",
			TableMap: notes,
		},
	})

	e.AddService(&UsageService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/usages",
			TableMap: usages,
		},
	})

	e.AddService(&WordService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/words",
			TableMap: words,
		},
	})

	e.AddService(&UserService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/users",
			TableMap: users,
		},
	})

	e.AddService(&LanguageService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/languages",
			TableMap: languages,
		},
	})

	e.AddService(&InstanceService{
		BaseService: BaseService{
			Engine:   e,
			Prefix:   "/api/instances",
			TableMap: instances,
		},
	})

	// Run

	e.Router.Run(e.Config.Address)

	return nil
}

// don't like this
// this will produce an error when it already exists but that's ok for now
func (e *Engine) AddForeignKey(table *gorp.TableMap, key string, reference *gorp.TableMap, column string, ordinal uint) error {
	if _, err := e.DbMap.Exec(fmt.Sprintf(
		"alter table %s add constraint %s foreign key (%s) references %s(%s);",
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
