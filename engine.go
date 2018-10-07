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
	DB     *gorp.DbMap
	Router *gin.Engine
}

func (e *Engine) Run(createTables, createIndexes, createForeignKeys bool) error {
	e.DB = &gorp.DbMap{}

	if e.Config.Database.LogMode {
		e.DB.TraceOn("[gorp]", log.New(os.Stdout, "langsuite:", log.Lmicroseconds))
	}

	switch e.Config.Database.Driver {
	case "postgres":
		e.DB.Dialect = gorp.PostgresDialect{}

	default:
		log.Fatal("sunpported database driver")

	}

	{
		db, err := sql.Open(e.Config.Database.Driver, e.Config.Database.Source)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to open %s database at %s: %s", e.Config.Database.Driver, e.Config.Database.Source, err.Error()))
		}
		defer db.Close()

		e.DB.Db = db
	}

	var (
		collection_tags       = AddTable(e.DB, CollectionTag{})
		collections           = AddTable(e.DB, Collection{})
		corpus_tags           = AddTable(e.DB, CorpusTag{})
		corpus_words          = AddTable(e.DB, CorpusWord{})
		corpora               = AddTable(e.DB, Corpus{})
		definition_link_types = AddTable(e.DB, DefinitionLinkType{})
		definition_links      = AddTable(e.DB, DefinitionLink{})
		definitions           = AddTable(e.DB, Definition{})
		highlights            = AddTable(e.DB, Highlight{})
		instances             = AddTable(e.DB, Instance{})
		languages             = AddTable(e.DB, Language{})
		lexica                = AddTable(e.DB, Lexica{})
		note_collections      = AddTable(e.DB, NoteCollection{})
		note_definitions      = AddTable(e.DB, NoteDefinition{})
		note_tags             = AddTable(e.DB, NoteTag{})
		notes                 = AddTable(e.DB, Note{})
		tags                  = AddTable(e.DB, Tag{})
		usages                = AddTable(e.DB, Usage{})
		users                 = AddTable(e.DB, User{})
		wordlist_items        = AddTable(e.DB, WordlistItem{})
		wordlists             = AddTable(e.DB, Wordlist{})
		words                 = AddTable(e.DB, Word{})
	)

	words.AddIndex("words_unique_idx", "Btree", []string{"word"}).SetUnique(true)

	if createTables {
		if err := e.DB.CreateTablesIfNotExists(); err != nil {
			return errors.New("failed to create tables: " + err.Error())
		}
	}

	if createIndexes {
		if err := e.DB.CreateIndex(); err != nil {
			log.Println("failed to create indexes: " + err.Error())
		}
	}

	if createForeignKeys {
		AddForeignKey(e.DB, collection_tags, "collection_id", collections, "id", 1)
		AddForeignKey(e.DB, collection_tags, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, collection_tags, "tag_id", tags, "id", 1)

		AddForeignKey(e.DB, collections, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, corpora, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, corpus_tags, "corpus_id", corpora, "id", 1)
		AddForeignKey(e.DB, corpus_tags, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, corpus_tags, "tag_id", tags, "id", 1)

		AddForeignKey(e.DB, corpus_words, "corpus_id", corpora, "id", 1)
		AddForeignKey(e.DB, corpus_words, "word_id", words, "id", 1)

		AddForeignKey(e.DB, definition_links, "definition1_id", definitions, "id", 1)
		AddForeignKey(e.DB, definition_links, "definition2_id", definitions, "id", 2)
		AddForeignKey(e.DB, definition_links, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, definition_links, "type_id", definition_link_types, "id", 1)

		AddForeignKey(e.DB, definitions, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, definitions, "word_id", words, "id", 1)

		AddForeignKey(e.DB, highlights, "corpus_id", corpora, "id", 1)
		AddForeignKey(e.DB, highlights, "corpus_word_id", corpus_words, "id", 1)
		AddForeignKey(e.DB, highlights, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, instances, "language_id", languages, "id", 1)
		AddForeignKey(e.DB, instances, "user_id", users, "id", 1)

		AddForeignKey(e.DB, lexica, "language_id", languages, "id", 1)

		AddForeignKey(e.DB, note_collections, "collection_id", collections, "id", 1)
		AddForeignKey(e.DB, note_collections, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, note_collections, "note_id", notes, "id", 1)
		AddForeignKey(e.DB, note_definitions, "definition_id", definitions, "id", 1)
		AddForeignKey(e.DB, note_definitions, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, note_definitions, "note_id", notes, "id", 1)

		AddForeignKey(e.DB, note_tags, "instance_id", instances, "id", 1)
		AddForeignKey(e.DB, note_tags, "note_id", notes, "id", 1)
		AddForeignKey(e.DB, note_tags, "tag_id", tags, "id", 1)

		AddForeignKey(e.DB, notes, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, tags, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, usages, "corpus_id", corpora, "id", 1)
		AddForeignKey(e.DB, usages, "definition_id", definitions, "id", 1)
		AddForeignKey(e.DB, usages, "instance_id", instances, "id", 1)

		AddForeignKey(e.DB, wordlist_items, "word_id", words, "id", 1)
		AddForeignKey(e.DB, wordlist_items, "wordlist_id", wordlists, "id", 1)
	}

	e.Router = gin.Default()
	e.Router.HTMLRender = multitemplate.NewRenderer()

	{
		v := VersionMiddleware{Engine: e}
		e.Router.Use(v.Handler)
	}

	(&WebService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/",
		},
	}).Register()

	(&WordService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + words.TableName,
		},
	}).Register()

	(&CorpusService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + corpora.TableName,
		},
	}).Register()

	(&DefinitionService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + definitions.TableName,
		},
	}).Register()

	(&InstanceService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + instances.TableName,
		},
	}).Register()

	(&LanguageService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + languages.TableName,
		},
	}).Register()

	(&NoteService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + notes.TableName,
		},
	}).Register()

	(&UsageService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + usages.TableName,
		},
	}).Register()

	(&UserService{
		BaseService: BaseService{
			Engine: e,
			Prefix: "/api/v1/" + users.TableName,
		},
	}).Register()

	e.Router.Run(e.Config.Address)

	return nil
}
