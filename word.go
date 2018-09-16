package langmap

type Word struct {
	Id   uint   `db:"id" json:"id"`
	Word string `db:"word" json:"word"`
}

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
