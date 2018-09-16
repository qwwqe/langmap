package langmap

type Note struct {
	Id         uint   `db:"id"`
	Title      string `db:"title"`
	Comment    string `db:"comment"`
	InstanceId uint   `db:"instance_id"`
}
