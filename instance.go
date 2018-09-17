package langmap

type Instance struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	UserId     uint   `db:"user_id"`
	LanguageId uint   `db:"language_id"`
}
