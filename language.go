package langmap

type Language struct {
	Id  uint   `db:"id"`
	Tag string `db:"tag"`
}
