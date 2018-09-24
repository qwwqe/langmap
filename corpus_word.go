package langmap

import (
	"database/sql"

	"github.com/go-gorp/gorp"
)

type CorpusWord struct {
	BaseTable
	Position uint    `json:"position" db:"position"`
	Sentence uint    `json:"sentence" db:"sentence"`
	Semantic bool    `json:"semantic" db:"semantic"`
	CorpusId uint    `json:"corpus_id" db:"corpus_id"`
	Corpus   *Corpus `json:"corpus,omitempty" db:"-"`
	WordId   uint    `json:"word_id" db:"word_id"`
	Word     *Word   `json:"word,omitempty" db:"-"`
}

func (CorpusWord) TableName() string { return "corpus_words" }

func (i *CorpusWord) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			i.Id = uint(v.(float64))

		case "position":
			i.Position = uint(v.(float64))

		case "sentence":
			i.Sentence = uint(v.(float64))

		case "semantic":
			i.Semantic = v.(bool)

		case "corpus_id":
			i.CorpusId = uint(v.(float64))

		case "word_id":
			i.WordId = uint(v.(float64))

		}
	}
}

func (t *CorpusWord) PreInsert(s gorp.SqlExecutor) error {
	if err := s.SelectOne(t.Word, "select id from "+t.Word.TableName()+" where word = $1", t.Word.Word); err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		if err := s.Insert(t.Word); err != nil {
			return err
		}
	}

	t.WordId = t.Word.Id

	return nil
}
