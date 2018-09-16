package langmap

type WordlistItem struct {
	Id         uint `db:"id"`
	WordlistId uint `db:"wordlist_id"`
	WordId     uint `db:"word_id"`
}
