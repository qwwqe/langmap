package langmap

type NoteCollection struct {
	BaseTable
	CollectionId uint `json:"collection_id" db:"collection_id"`
	NoteId       uint `json:"note_id" db:"note_id"`
	InstanceId   uint `json:"instance_id" db:"instance_id"`
}
