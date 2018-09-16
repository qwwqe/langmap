package langmap

type NoteTag struct {
	Id         uint `db:"id"`
	NoteId     uint `db:"note_id"`
	TagId      uint `db:"tag_id"`
	InstanceId uint `db:"instance_id"`
}
