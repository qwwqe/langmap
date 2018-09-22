package langmap

type Lexica struct {
	BaseTable
	URI        string `db:"uri"`
	Name       string `db:"name"`
	LanguageId uint   `db:"language_id"`
}
