package langmap

type Language struct {
	BaseTable
	Tag string `json:"tag" db:"tag"`
}

func (Language) TableName() string { return "languages" }

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
