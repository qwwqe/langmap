package langmap

type Usage struct {
	BaseTable
	DefinitionId uint `json:"definition_id" db:"definition_id"`
	CorpusId     uint `json:"corpus_id" db:"corpus_id"`
	InstanceId   uint `json:"instance_id" db:"instance_id"`
}

func (_ Usage) TableName() string { return "usages" }

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
