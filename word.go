package langmap

import "github.com/go-gorp/gorp"

type Word struct {
	BaseTable
	Word string `json:"word" db:"word"`
}

func (Word) TableName() string { return "words" }

func LoadWords(db *gorp.DbMap, f Filter) ([]*Word, error) {
	r := []*Word{}

	if _, err := db.Select(&r, SelectQuery(Word{}, f), f.Values...); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

func (i *Word) Inject(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			i.Id = uint(v.(float64))

		case "word":
			i.Word = v.(string)

		}
	}
}
