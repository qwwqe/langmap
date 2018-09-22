package langmap

type WordlistItem struct {
	BaseTable
	WordlistId uint `json:"wordlist_id" db:"wordlist_id"`
	WordId     uint `json:"word_id" db:"word_id"`
}
