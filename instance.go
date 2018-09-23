package langmap

type Instance struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	UserId     uint   `json:"user_id" db:"user_id"`
	LanguageId uint   `json:"language_id" db:"language_id"`
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
