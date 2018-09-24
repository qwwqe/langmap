package langmap

import "github.com/go-gorp/gorp"

type Instance struct {
	BaseTable
	Name       string    `json:"name" db:"name"`
	UserId     uint      `json:"user_id" db:"user_id"`
	LanguageId uint      `json:"language_id" db:"language_id"`
	Language   *Language `json:"language,omitempty" db:"-"`
}

func (Instance) TableName() string { return "instances" }

func (i *Instance) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			i.Id = uint(v.(float64))

		case "name":
			i.Name = v.(string)

		case "user_id":
			i.UserId = uint(v.(float64))

		case "language_id":
			i.LanguageId = uint(v.(float64))

		}
	}
}

func (i *Instance) Preload(db *gorp.DbMap) error {
	l := &Language{}

	if err := db.SelectOne(l, "select * from "+l.TableName()+" where id = $1", i.LanguageId); err != nil {
		return err
	}

	i.Language = l

	return nil
}
