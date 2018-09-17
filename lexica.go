package langmap

type Lexica struct {
	Id         uint   `db:"id"`
	URI        string `db:"uri"`
	Name       string `db:"name"`
	LanguageId uint   `db:"language_id"`
}
