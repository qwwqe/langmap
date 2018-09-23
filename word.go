package langmap

type Word struct {
	BaseTable
	Word string `json:"word" db:"word"`
}

func (Word) TableName() string { return "words" }

func (w *Word) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			w.Id = uint(v.(float64))

		case "word":
			w.Word = v.(string)

		}
	}
}
