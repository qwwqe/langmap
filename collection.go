package langmap

type Collection struct {
	BaseTable
	Name       string `db:"name" json:"name"`
	InstanceId uint   `db:"instance_id" json:"instance_id"`
}
