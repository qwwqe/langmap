package langmap

type CorpusWord struct {
	BaseTable
	Position uint `json:"position" db:"position"`
	Sentence uint `json:"sentence" db:"sentence"`
	Semantic bool `json:"semantic" db:"semantic"`
	CorpusId uint `json:"corpus_id" db:"corpus_id"`
	WordId   uint `json:"word_id" db:"word_id"`
}
