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

	Word string `json:"word" db:"-"`
}

func (d *Definition) PreInsert(s gorp.SqlExecutor) error {
	w := Word{}

	if err := s.SelectOne(&w, "select * from words where word = $1", d.Word); err != nil {
		if err == sql.ErrNoRows {
			w.Word = d.Word

			if err := s.Insert(&w); err != nil {
				return err
			}

		} else {
			return err

		}
	}

	d.WordId = w.Id

	return nil
}

func (d *Definition) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			d.Id = uint(v.(float64))

		case "pronunciation":
			d.Pronunciation = v.(string)

		case "meaning":
			d.Meaning = v.(string)

		case "word_id":
			d.WordId = uint(v.(float64))

		case "instance_id":
			d.InstanceId = uint(v.(float64))

		}
	}
}
