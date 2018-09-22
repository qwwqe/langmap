package langmap

type WordlistItem struct {
	BaseTable
	WordlistId uint `db:"wordlist_id"`
	WordId     uint `db:"word_id"`
}
