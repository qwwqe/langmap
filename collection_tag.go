package langmap

type CollectionTag struct {
	BaseTable
	CollectionId uint `json:"collection_id" db:"collection_id"`
	TagId        uint `json:"tag_id" db:"tag_id"`
	InstanceId   uint `json:"instance_id" db:"instance_id"`
}

func (CollectionTag) TableName() string { return "collection_tags" }
