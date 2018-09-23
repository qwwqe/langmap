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
	s.SetEngine(e)
	s.Register()
}

func (e *Engine) Run(createTables, createIndexes, createForeignKeys bool) error {

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
		collection_tags       = e.AddTable(CollectionTag{})
		collections           = e.AddTable(Collection{})
		corpus_tags           = e.AddTable(CorpusTag{})
		corpus_words          = e.AddTable(CorpusWord{})
		corpora               = e.AddTable(Corpus{})
		definition_link_types = e.AddTable(DefinitionLinkType{})
		definition_links      = e.AddTable(DefinitionLink{})
		definitions           = e.AddTable(Definition{})
		highlights            = e.AddTable(Highlight{})
		instances             = e.AddTable(Instance{})
		languages             = e.AddTable(Language{})
		lexica                = e.AddTable(Lexica{})
		note_collections      = e.AddTable(NoteCollection{})
		note_definitions      = e.AddTable(NoteDefinition{})
		note_tags             = e.AddTable(NoteTag{})
		notes                 = e.AddTable(Note{})
		tags                  = e.AddTable(Tag{})
		usages                = e.AddTable(Usage{})
		users                 = e.AddTable(User{})
		wordlist_items        = e.AddTable(WordlistItem{})
		wordlists             = e.AddTable(Wordlist{})
		words                 = e.AddTable(Word{})
	)

	words.AddIndex("words_unique_idx", "Btree", []string{"word"}).SetUnique(true)

	if createTables {
		if err := e.DbMap.CreateTablesIfNotExists(); err != nil {
			return errors.New("failed to create tables: " + err.Error())
		}
	}

	if createIndexes {
		if err := e.DbMap.CreateIndex(); err != nil {
			log.Println("failed to create indexes: " + err.Error())
		}
	}

	if createForeignKeys {
		for _, t := range []*gorp.TableMap{
			wordlist_items,
			corpus_words,
			definitions,
			highlights,
		} {
			e.AddForeignKey(t, "word_id", words, "id", 1)
		}

		for _, t := range []*gorp.TableMap{
			collection_tags,
			collections,
			corpora,
			corpus_tags,
			definition_links,
			definitions,
			highlights,
			note_collections,
			note_definitions,
			note_tags,
			notes,
			tags,
			usages,
			wordlists,
		} {
			e.AddForeignKey(t, "instance_id", instances, "id", 1)
		}

		e.AddForeignKey(collection_tags, "collection_id", collections, "id", 1)
		e.AddForeignKey(collection_tags, "tag_id", tags, "id", 1)
		e.AddForeignKey(corpus_tags, "corpus_id", corpora, "id", 1)
		e.AddForeignKey(corpus_tags, "tag_id", tags, "id", 1)
		e.AddForeignKey(corpus_words, "corpus_id", corpora, "id", 1)
		e.AddForeignKey(definition_links, "definition1_id", definitions, "id", 1)
		e.AddForeignKey(definition_links, "definition2_id", definitions, "id", 2)
		e.AddForeignKey(definition_links, "type_id", definition_link_types, "id", 1)
		e.AddForeignKey(highlights, "corpus_id", corpora, "id", 1)
		e.AddForeignKey(instances, "language_id", languages, "id", 1)
		e.AddForeignKey(instances, "user_id", users, "id", 1)
		e.AddForeignKey(lexica, "language_id", languages, "id", 1)
		e.AddForeignKey(note_collections, "collection_id", collections, "id", 1)
		e.AddForeignKey(note_collections, "note_id", notes, "id", 1)
		e.AddForeignKey(note_definitions, "definition_id", definitions, "id", 1)
		e.AddForeignKey(note_definitions, "note_id", notes, "id", 1)
		e.AddForeignKey(note_tags, "note_id", notes, "id", 1)
		e.AddForeignKey(note_tags, "tag_id", tags, "id", 1)
		e.AddForeignKey(usages, "corpus_id", corpora, "id", 1)
		e.AddForeignKey(usages, "definition_id", definitions, "id", 1)
		e.AddForeignKey(wordlist_items, "wordlist_id", wordlists, "id", 1)
	}

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

	for _, s := range []RoutableService{
		&DefinitionService{BaseService: BaseService{Prefix: "/api/definitions"}},
		&NoteService{BaseService: BaseService{Prefix: "/api/notes"}},
		&UsageService{BaseService: BaseService{Prefix: "/api/usages"}},
		&WordService{BaseService: BaseService{Prefix: "/api/words"}},
		&UserService{BaseService: BaseService{Prefix: "/api/users"}},
		&LanguageService{BaseService: BaseService{Prefix: "/api/languages"}},
		&InstanceService{BaseService: BaseService{Prefix: "/api/instances"}},
	} {
		e.AddService(s)
	}

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

func (e *Engine) AddTable(i IdentifiableTable) *gorp.TableMap {
	return e.DbMap.AddTableWithName(i, i.TableName()).SetKeys(true, "id")
}
