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

func LoadDefinitions(db *gorp.DbMap, id uint) ([]*Definition, error) {
	r := []*Definition{}
	q := "select * from " + Definition{}.TableName() + " where instance_id = $1"

	if _, err := db.Select(&r, q, id); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

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

func (r *Definition) Preload(db *gorp.DbMap) error {
	r.Word = &Word{}
	LoadOne(db, r.Word, r.WordId)

	return nil
}
