package langmap

type NoteTag struct {
	BaseTable
	NoteId     uint `db:"note_id"`
	TagId      uint `db:"tag_id"`
	InstanceId uint `db:"instance_id"`
}
