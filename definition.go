package langmap

import (
	"database/sql"

	"github.com/go-gorp/gorp"
)

type Definition struct {
	Id            uint   `json:"id" db:"id"`
	Pronunciation string `json:"pronunciation" db:"pronunciation"`
	Meaning       string `json:"meaning" db:"meaning"`
	WordId        uint   `json:"word_id" db:"word_id"`
	InstanceId    uint   `json:"instance_id" db:"instance_id"`

	Word string `json:"word" db:"-"`
}

// (d *Definition) PreInsert retrieves or creates a Word before the
// creation of a Definition and fills in that Definition's corresponding
// WordId field
func (d *Definition) PreInsert(s gorp.SqlExecutor) error {
	var word Word
	err := s.SelectOne(&word, "select * from words where word=?", d.Word)

	switch err {
	case nil: // Word exists, simply update WordId
		d.WordId = word.Id
	case sql.ErrNoRows: // Word does not exist, create it first
		word = Word{Word: d.Word}
		err := s.Insert(&word)
		if err != nil {
			return err
		}
		d.WordId = word.Id
	default:
		return err
	}

	return nil
}

func (d *Definition) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			d.Id = v.(uint)

		case "pronunciation":
			d.Pronunciation = v.(string)

		case "meaning":
			d.Meaning = v.(string)

		case "word_id":
			d.WordId = v.(uint)

		case "instance_id":
			d.InstanceId = v.(uint)

		}
	}
}
