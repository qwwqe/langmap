package langmap

type Tag struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (Tag) TableName() string { return "tags" }
