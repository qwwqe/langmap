package langmap

type NoteCollection struct {
	BaseTable
	CollectionId uint `db:"collection_id"`
	NoteId       uint `db:"note_id"`
	InstanceId   uint `db:"instance_id"`
}
