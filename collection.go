package langmap

type Collection struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (Collection) TableName() string { return "collections" }
