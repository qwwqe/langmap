package langmap

type NoteCollection struct {
	Id           uint `db:"id"`
	CollectionId uint `db:"collection_id"`
	NoteId       uint `db:"note_id"`
	InstanceId   uint `db:"instance_id"`
}
