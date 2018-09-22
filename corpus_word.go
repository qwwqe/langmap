package langmap

type CorpusWord struct {
	BaseTable
	Position uint `db:"position"`
	Sentence uint `db:"sentence"`
	Semantic bool `db:"semantic"`
	CorpusId uint `db:"corpus_id"`
	WordId   uint `db:"word_id"`
}
