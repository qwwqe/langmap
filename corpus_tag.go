package langmap

type CorpusTag struct {
	BaseTable
	CorpusId   uint `json:"corpus_id" db:"corpus_id"`
	TagId      uint `json:"tag_id" db:"tag_id"`
	InstanceId uint `json:"instance_id" db:"instance_id"`
}

func (_ CorpusTag) TableName() string { return "corpus_tags" }
