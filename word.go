package langmap

type Word struct {
	Id   uint   `db:"id"`
	Word string `db:"word"`
}

func (w *Word) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			w.Id = v.(uint)

		case "word":
			w.Word = v.(string)

		}
	}
}
