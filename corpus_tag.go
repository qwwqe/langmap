package langmap

type CorpusTag struct {
	Id         uint `db:"id"`
	CorpusId   uint `db:"corpus_id"`
	TagId      uint `db:"tag_id"`
	InstanceId uint `db:"instance_id"`
}
