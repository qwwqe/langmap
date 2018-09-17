package langmap

type User struct {
	Id   uint   `db:"id"`
	Name string `db:"name"`
}

func (u *User) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			u.Id = uint(v.(float64))

		case "name":
			u.Name = v.(string)

		}
	}
}
