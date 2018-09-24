package langmap

type User struct {
	BaseTable
	Name string `json:"name" db:"name"`
}

func (User) TableName() string { return "users" }

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
