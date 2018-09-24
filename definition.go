package langmap

import (
	"database/sql"

	"github.com/go-gorp/gorp"
)

type Definition struct {
	BaseTable
	Pronunciation string `json:"pronunciation" db:"pronunciation"`
	Meaning       string `json:"meaning" db:"meaning"`
	WordId        uint   `json:"word_id" db:"word_id"`
	InstanceId    uint   `json:"instance_id" db:"instance_id"`
	Word          *Word  `json:"word,omitempty" db:"-"`
}

func (Definition) TableName() string { return "definitions" }

func (i *Definition) Inject(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			i.Id = uint(v.(float64))

		case "pronunciation":
			i.Pronunciation = v.(string)

		case "meaning":
			i.Meaning = v.(string)

		case "word_id":
			i.WordId = uint(v.(float64))

		case "instance_id":
			i.InstanceId = uint(v.(float64))

		}
	}
}

func (t *Definition) PreInsert(s gorp.SqlExecutor) error {
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

func (t *Definition) Preload(db *gorp.DbMap) error {
	w := &Word{}

	if err := db.SelectOne(w, "select * from "+w.TableName()+" where id = $1", t.WordId); err != nil {
		return err
	}

	t.Word = w

	return nil
}
