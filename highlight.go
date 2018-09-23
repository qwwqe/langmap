package langmap

type Highlight struct {
	BaseTable
	CorpusId     uint `json:"corpus_id" db:"corpus_id"`
	CorpusWordId uint `json:"corpus_word_id" db:"corpus_word_id"`
	InstanceId   uint `json:"instance_id" db:"instance_id"`
}

func (Highlight) TableName() string { return "highlights" }
