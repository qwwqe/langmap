package langmap

type Tag struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	InstanceId uint   `db:"instance_id"`
}
