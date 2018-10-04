package langmap

import "github.com/go-gorp/gorp"

type User struct {
	BaseTable
	Name string `json:"name" db:"name"`
}

func (User) TableName() string { return "users" }

func LoadUsers(db *gorp.DbMap) ([]*User, error) {
	r := []*User{}
	q := "select * from " + User{}.TableName()

	if _, err := db.Select(&r, q); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

func (u *User) Inject(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			u.Id = uint(v.(float64))

		case "name":
			u.Name = v.(string)

		}
	}
}
