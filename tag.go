package langmap

type Tag struct {
	BaseTable
	Name       string `db:"name"`
	InstanceId uint   `db:"instance_id"`
}
