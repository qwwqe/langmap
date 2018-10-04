package langmap

import "github.com/go-gorp/gorp"

type Language struct {
	BaseTable
	Tag string `json:"tag" db:"tag"`
}

func (Language) TableName() string { return "languages" }

func LoadLanguages(db *gorp.DbMap) ([]*Language, error) {
	r := []*Language{}
	q := "select * from " + Language{}.TableName()

	if _, err := db.Select(&r, q); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

func (l *Language) Inject(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			l.Id = uint(v.(float64))

		case "tag":
			l.Tag = v.(string)

		}
	}
}
