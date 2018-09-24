package langmap

type Word struct {
	BaseTable
	Word string `json:"word" db:"word"`
}

func (Word) TableName() string { return "words" }

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
