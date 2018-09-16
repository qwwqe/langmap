package langmap

type Usage struct {
	Id           uint `db:"id"`
	DefinitionId uint `db:"definition_id"`
	CorpusId     uint `db:"corpus_id"`
	InstanceId   uint `db:"instance_id"`
}

func (u *Usage) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			u.Id = v.(uint)

		case "definition_id":
			u.DefinitionId = v.(uint)

		case "corpus_id":
			u.CorpusId = v.(uint)

		case "instance_id":
			u.InstanceId = v.(uint)

		}
	}
}
