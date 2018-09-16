package langmap

type Highlight struct {
	Id           uint `db:"id"`
	CorpusId     uint `db:"corpus_id"`
	CorpusWordId uint `db:"corpus_word_id"`
	InstanceId   uint `db:"instance_id"`
}
