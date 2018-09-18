package langmap

type Instance struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	UserId     uint   `db:"user_id"`
	LanguageId uint   `db:"language_id"`
}

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
