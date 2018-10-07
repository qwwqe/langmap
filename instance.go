package langmap

import "github.com/go-gorp/gorp"

type Instance struct {
	BaseTable
	Name       string    `json:"name" db:"name"`
	UserId     uint      `json:"user_id" db:"user_id"`
	User       *User     `json:"user,omitempty" db:"-"`
	LanguageId uint      `json:"language_id" db:"language_id"`
	Language   *Language `json:"language,omitempty" db:"-"`
}

func (Instance) TableName() string { return "instances" }

func LoadInstances(db *gorp.DbMap, f Filter) ([]*Instance, error) {
	r := []*Instance{}

	if _, err := db.Select(&r, SelectQuery(Instance{}, f), f.Values...); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

func (i *Instance) Inject(m map[string]interface{}) {
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

func (r *Instance) Preload(db *gorp.DbMap) error {
	r.User = &User{}
	LoadOne(db, r.User, r.UserId)

	r.Language = &Language{}
	LoadOne(db, r.Language, r.LanguageId)

	return nil
}
