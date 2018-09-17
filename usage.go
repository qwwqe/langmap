package langmap

type Usage struct {
	Id           uint `db:"id" json:"id"`
	DefinitionId uint `db:"definition_id" json:"definition_id"`
	CorpusId     uint `db:"corpus_id" json:"corpus_id"`
	InstanceId   uint `db:"instance_id" json:"instance_id"`
}

func (u *Usage) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			u.Id = uint(v.(float64))

		case "definition_id":
			u.DefinitionId = uint(v.(float64))

		case "corpus_id":
			u.CorpusId = uint(v.(float64))

		case "instance_id":
			u.InstanceId = uint(v.(float64))

		}
	}
}
