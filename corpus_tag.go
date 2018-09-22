package langmap

type CorpusTag struct {
	BaseTable
	CorpusId   uint `db:"corpus_id"`
	TagId      uint `db:"tag_id"`
	InstanceId uint `db:"instance_id"`
}
