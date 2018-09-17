package langmap

type CollectionTag struct {
	Id           uint `db:"id"`
	CollectionId uint `db:"collection_id"`
	TagId        uint `db:"tag_id"`
	InstanceId   uint `db:"instance_id"`
}
